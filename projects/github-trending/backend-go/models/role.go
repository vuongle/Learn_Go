package models

// Define a enum "Role"
type Role int

const (
	MEMBER Role = iota
	ADMIN
)

// Convert role from int to string
func (r Role) String() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
