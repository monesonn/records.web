package models

// SignUp struct to describe register a new user.
type SignUp struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Username string `json:"username" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"category" validate:"lte=25"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
