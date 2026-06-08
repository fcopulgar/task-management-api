package persistence

import "time"

type UserModel struct {
	ID                 string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name               string `gorm:"not null"`
	Email              string `gorm:"uniqueIndex;not null"`
	PasswordHash       string `gorm:"not null"`
	Role               string `gorm:"not null"`
	MustChangePassword bool   `gorm:"not null;default:true"`
	Active             bool   `gorm:"not null;default:true"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (UserModel) TableName() string {
	return "users"
}

type SessionModel struct {
	ID        string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string     `gorm:"not null;index"`
	RevokedAt *time.Time
	ExpiresAt time.Time  `gorm:"not null"`
	CreatedAt time.Time
}

func (SessionModel) TableName() string {
	return "sessions"
}

type TaskModel struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string    `gorm:"not null"`
	Description string
	DueAt       time.Time `gorm:"not null"`
	Status      string    `gorm:"not null;default:ASSIGNED"`
	AssigneeID  string    `gorm:"not null;index"`
	CreatedBy   string    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TaskModel) TableName() string {
	return "tasks"
}

type CommentModel struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	TaskID    string    `gorm:"not null;index"`
	UserID    string    `gorm:"not null"`
	Comment   string    `gorm:"not null"`
	CreatedAt time.Time
}

func (CommentModel) TableName() string {
	return "task_comments"
}
