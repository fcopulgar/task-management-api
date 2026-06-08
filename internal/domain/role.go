package domain

type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleExecutor Role = "EXECUTOR"
	RoleAuditor  Role = "AUDITOR"
)

func (r Role) IsValid() bool {
	return r == RoleAdmin || r == RoleExecutor || r == RoleAuditor
}

func (r Role) String() string {
	return string(r)
}
