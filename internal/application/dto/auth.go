package dto

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string `json:"token"`
}

type ChangePasswordInput struct {
	UserID      string
	OldPassword string
	NewPassword string
}
