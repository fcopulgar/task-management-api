package ports

import (
	"testing"
)

func TestInterfaceCompliance(t *testing.T) {
	var _ UserRepository = (*MockUserRepository)(nil)
	var _ SessionRepository = (*MockSessionRepository)(nil)
	var _ TaskRepository = (*MockTaskRepository)(nil)
	var _ CommentRepository = (*MockCommentRepository)(nil)
	var _ PasswordHasher = (*MockPasswordHasher)(nil)
	var _ TokenService = (*MockTokenService)(nil)

	t.Log("todas las interfaces compilan con sus mocks")
}
