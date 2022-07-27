package rbac

import (
	"context"
	_ "embed"
	"fmt"
	"golang.org/x/xerrors"

	"github.com/open-policy-agent/opa/rego"
)

type Authorizer interface {
	ByRoleName(ctx context.Context, subjectID string, roleNames []string, action Action, object Object) error
}

// Filter takes in a list of objects, and will filter the list removing all
// the elements the subject does not have permission for.
// Filter does not allocate a new slice, and will use the existing one
// passed in. This can cause memory leaks if the slice is held for a prolonged
// period of time.
func Filter[O Objecter](ctx context.Context, auth Authorizer, subjID string, subjRoles []string, action Action, objects []O) []O {
	filtered := make([]O, 0)

	for i := range objects {
		object := objects[i]
		err := auth.ByRoleName(ctx, subjID, subjRoles, action, object.RBACObject())
		if err == nil {
			filtered = append(filtered, object)
		}
	}
	return filtered
}

// RegoAuthorizer will use a prepared rego query for performing authorize()
type RegoAuthorizer struct {
	query rego.PreparedEvalQuery
}

// Load the policy from policy.rego in this directory.
//go:embed policy.rego
var policy string

func NewAuthorizer() (*RegoAuthorizer, error) {
	ctx := context.Background()
	query, err := rego.New(
		// Query returns true/false for authorization access
		rego.Query("data.authz.allow"),
		rego.Module("policy.rego", policy),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, xerrors.Errorf("prepare query: %w", err)
	}
	return &RegoAuthorizer{query: query}, nil
}

type authSubject struct {
	ID    string `json:"id"`
	Roles []Role `json:"roles"`
}

// ByRoleName will expand all roleNames into roles before calling Authorize().
// This is the function intended to be used outside this package.
// The role is fetched from the builtin map located in memory.
func (a RegoAuthorizer) ByRoleName(ctx context.Context, subjectID string, roleNames []string, action Action, object Object) error {
	roles := make([]Role, 0, len(roleNames))
	for _, n := range roleNames {
		r, err := RoleByName(n)
		if err != nil {
			return xerrors.Errorf("get role permissions: %w", err)
		}
		roles = append(roles, r)
	}
	return a.Authorize(ctx, subjectID, roles, action, object)
}

// Authorize allows passing in custom Roles.
// This is really helpful for unit testing, as we can create custom roles to exercise edge cases.
func (a RegoAuthorizer) Authorize(ctx context.Context, subjectID string, roles []Role, action Action, object Object) error {
	input := map[string]interface{}{
		"subject": authSubject{
			ID:    subjectID,
			Roles: roles,
		},
		"object": object,
		"action": action,
	}

	results, err := a.query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return ForbiddenWithInternal(xerrors.Errorf("eval rego: %w", err), input, results)
	}

	if !results.Allowed() {
		return ForbiddenWithInternal(xerrors.Errorf("policy disallows request"), input, results)
	}

	return nil
}

// Load the policy from policy.rego in this directory.
//go:embed partial.rego
var partial string

func (a RegoAuthorizer) Partial(ctx context.Context, subjectID string, roleNames []string, action Action, object Object) error {
	roles := make([]Role, 0, len(roleNames))
	for _, n := range roleNames {
		r, err := RoleByName(n)
		if err != nil {
			return xerrors.Errorf("get role permissions: %w", err)
		}
		roles = append(roles, r)
	}

	input := map[string]interface{}{
		"subject": authSubject{
			ID:    subjectID,
			Roles: roles,
		},
		"object": map[string]string{
			"type": object.Type,
		},
		"action": action,
	}

	query, err := rego.New(
		// Query returns true/false for authorization access
		rego.Query("data.authz.allow = true"),
		rego.Module("partial.rego", partial),
		rego.Unknowns([]string{
			"input.object.owner",
			"input.object.org_owner",
		}),
	).PrepareForPartial(ctx)

	part, err := query.Partial(ctx, rego.EvalInput(input))
	if err != nil {
		return ForbiddenWithInternal(xerrors.Errorf("eval rego: %w", err), input, nil)
	}

	for _, q := range part.Queries {
		fmt.Println(q.String())
	}
	fmt.Println("--")
	for _, s := range part.Support {
		fmt.Println(s.String())
	}
	fmt.Println("---")

	// REGULAR
	input["object"] = object
	rs, err := rego.New(
		// Query returns true/false for authorization access
		rego.Query("data.authz.allow"),
		rego.Module("partial.rego", partial),
		rego.Input(input),
	).Eval(ctx)
	fmt.Println("Allow:", rs.Allowed())

	return nil
}
