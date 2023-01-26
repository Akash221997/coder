package rbac

// Subject is a struct that contains all the elements of a subject in an rbac
// authorize.
type Subject struct {
	SubjectID string
	Roles     ExpandableRoles
	Groups    []string
	Scope     ScopeName
}
