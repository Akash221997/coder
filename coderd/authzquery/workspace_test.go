package authzquery_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/coder/coder/coderd/rbac"
	"github.com/moby/moby/pkg/namesgenerator"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coder/coder/coderd/database"

	"github.com/coder/coder/coderd/authzquery"
	"github.com/coder/coder/coderd/coderdtest"
	"github.com/coder/coder/coderd/database/databasefake"
)

func DBSetup(t *testing.T, opts ...func(database.Store) error) database.Store {
	t.Helper()

	db := databasefake.New()
	for _, opt := range opts {
		require.NoError(t, opt(db))
	}

	return db
}

type authzTestCase[Shared any] struct {
	Setup func(*testing.T, database.Store) Shared
	RunF  func(*testing.T, Shared, authzquery.AuthzStore)
}

//func TestWorkspaceFunctions(t *testing.T) {
//
//	testCases := []struct {
//		Setup func(context.Context, database.Store) error
//		RunF  func(context.Context, authzquery.AuthzStore) error
//	}{
//		{
//			Setup: func(ctx context.Context, store database.Store) error {
//				workspace := (&testWorkspace{db: store}).Insert(t)
//				return nil
//			},
//			RunF: func(ctx context.Context, store authzquery.AuthzStore) error {
//				_, err := store.GetWorkspaceByID(ctx, uuid.New())
//				return err
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		//t.Run(tc.name, func(t *testing.T) {
//		//
//		//})
//	}
//}

func TestWorkspace(t *testing.T) {
	// GetWorkspaceByID
	var (
		db = databasefake.New()
		// TODO: Recorder should record all authz calls
		rec   = &coderdtest.RecordingAuthorizer{}
		q     = authzquery.NewAuthzQuerier(db, rec)
		actor = rbac.Subject{
			SubjectID: uuid.New().String(),
			Roles:     rbac.RoleNames{rbac.RoleOwner()},
			Groups:    []string{},
			Scope:     rbac.ScopeAll,
		}
		ctx = authzquery.WithAuthorizeSubjectContext(context.Background(), actor)
	)

	workspace := insertRandomWorkspace(t, db)

	// Test recorder
	_, err := q.GetWorkspaceByID(ctx, workspace.ID)
	require.NoError(t, err)

	_, err = q.UpdateWorkspace(ctx, database.UpdateWorkspaceParams{
		ID:   workspace.ID,
		Name: "new-name",
	})
	require.NoError(t, err)

	rec.AssertActor(t, actor,
		rec.Pair(rbac.ActionRead, workspace),
		rec.Pair(rbac.ActionUpdate, workspace),
	)
	require.NoError(t, rec.AllAsserted())
}

func insertRandomWorkspace(t *testing.T, db database.Store, opts ...func(w *database.Workspace)) database.Workspace {
	workspace := &database.Workspace{
		ID:             uuid.New(),
		CreatedAt:      time.Now().Add(time.Hour * -1),
		UpdatedAt:      time.Now(),
		OwnerID:        uuid.New(),
		OrganizationID: uuid.New(),
		TemplateID:     uuid.New(),
		Deleted:        false,
		Name:           namesgenerator.GetRandomName(1),
		LastUsedAt:     time.Now(),
	}
	for _, opt := range opts {
		opt(workspace)
	}

	newWorkspace, err := db.InsertWorkspace(context.Background(), database.InsertWorkspaceParams{
		ID:                workspace.ID,
		CreatedAt:         workspace.CreatedAt,
		UpdatedAt:         workspace.UpdatedAt,
		OwnerID:           workspace.OwnerID,
		OrganizationID:    workspace.OrganizationID,
		TemplateID:        workspace.TemplateID,
		Name:              workspace.Name,
		AutostartSchedule: workspace.AutostartSchedule,
		Ttl:               workspace.Ttl,
	})
	require.NoError(t, err, "insert workspace")
	return newWorkspace
}

func TestRandomStruct(t *testing.T) {
	var w database.Workspace
	err := faker.FakeData(&w)
	require.NoError(t, err)
}

//func newWorkspace(opts ...func(w *database.Workspace)) {
//	w := &database.Workspace{
//		ID:                uuid.UUID{},
//		CreatedAt:         time.Time{},
//		UpdatedAt:         time.Time{},
//		OwnerID:           uuid.UUID{},
//		OrganizationID:    uuid.UUID{},
//		TemplateID:        uuid.UUID{},
//		Deleted:           false,
//		Name:              "",
//		AutostartSchedule: sql.NullString{},
//		Ttl:               sql.NullInt64{},
//		LastUsedAt:        time.Time{},
//	}
//	for _, opt := range opts {
//
//	}
//}

type testWorkspace struct {
	Workspace *database.Workspace
	db        database.Store
}

type testUser struct {
	User *database.User
	db   database.Store
}

func (u *testUser) Insert(t *testing.T) *testUser {
	u.User = gen(t, u.User)
	user, err := u.db.InsertUser(context.Background(), database.InsertUserParams{
		ID:             u.User.ID,
		Email:          u.User.Email,
		Username:       u.User.Username,
		HashedPassword: u.User.HashedPassword,
		CreatedAt:      u.User.CreatedAt,
		UpdatedAt:      u.User.UpdatedAt,
		RBACRoles:      u.User.RBACRoles,
		LoginType:      u.User.LoginType,
	})
	require.NoError(t, err, "failed to insert user")
	u.User = &user

	return u
}

func (w *testWorkspace) User(t *testing.T) *testUser {
	u := gen(t, &database.User{})
	u.ID = w.Workspace.OwnerID
	return &testUser{
		User: u,
		db:   w.db,
	}
}

func (w *testWorkspace) Insert(t *testing.T) *testWorkspace {
	w.Workspace = gen(t, w.Workspace)
	workspace, err := w.db.InsertWorkspace(context.Background(), database.InsertWorkspaceParams{
		ID:                w.Workspace.ID,
		CreatedAt:         w.Workspace.CreatedAt,
		UpdatedAt:         w.Workspace.UpdatedAt,
		OwnerID:           w.Workspace.OwnerID,
		OrganizationID:    w.Workspace.OrganizationID,
		TemplateID:        w.Workspace.TemplateID,
		Name:              w.Workspace.Name,
		AutostartSchedule: w.Workspace.AutostartSchedule,
		Ttl:               w.Workspace.Ttl,
	})
	require.NoError(t, err, "insert workspace")
	w.Workspace = &workspace
	return nil
}

func gen[Object any](t *testing.T, v *Object) *Object {
	if v == nil {
		err := faker.FakeData(v)
		require.NoError(t, err, "failed to generate fake %s", reflect.TypeOf(v).String())
	}
	return v
}
