package dto

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

type ChangePasswordInput struct {
	UserID      string `json:"-"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
