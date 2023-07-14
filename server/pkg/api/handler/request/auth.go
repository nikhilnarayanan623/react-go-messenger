package request

type Login struct {
	UserName string `json:"user_name" binding:"omitempty,min=3,max=15"`
	Phone    string `json:"phone" binding:"omitempty,min=10,max=10"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"required,min=5,max=30"`
}

type GoogleLogin struct {
	TokenCode string `json:"token" binding:"required"`
}

type UserSignUp struct {
	UserName        string `json:"user_name"  binding:"required,min=3,max=15"`
	FirstName       string `json:"first_name"  binding:"required,min=2,max=50"`
	LastName        string `json:"last_name"  binding:"required,min=1,max=50"`
	Age             uint   `json:"age"  binding:"required,numeric"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"min=10"`
}
