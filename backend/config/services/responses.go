package services

// Request
type RegisterRequest struct {
	Name     string `json:"name" example:"User's name"`
	Email    string `example:"test@email.com"`
	Password string `example:"12345"`
}

type LoginRequest struct {
	Email    string `example:"test@email.com"`
	Password string `example:"12345"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" example:"User's name"`
	Email string `json:"email" example:"User's email"`
}

type UpadatePasswordRequest struct {
	CurrentPassword    string `json:"current_password" example:"User's Current Password"`
	NewPassword        string `json:"new_password" example:"User's New Password"`
	NewPasswordConfirm string `json:"new_password_confirm" example:"User's Confirm Password"`
}

// succes response
type UserResponse struct {
	ID    string `json:"id" example:"User's ID"`
	Name  string `json:"name" example:"User's name"`
	Email string `json:"email" example:"User's email"`
}

type CurrentUserStatusResponse struct {
	IsLock bool `json:"is_lock" example:"false"`
}

type ChangePasswordResponse struct {
	Massage     string `json:"message" example:"Password updated successfully"`
	NewPassword string `json:"new_password" example:"su3cl3"`
}
type SuccessfulLoginResponse struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRC"`
	Message string `json:"message" example:"User login successfully"`
}

type SuccessfulSignup struct {
	Message string `json:"message" example:"User registered successfully"`
}

type SuccessfulLogoutResponse struct {
	Message string `json:"message" example:"Logged out successfully"`
}

type CreateSuccessfully struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"someone created successful"`
}

// error response
type BadRequest struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"Incorrect email"`
}

type HttpError struct {
	Code  int    `json:"code" example:"401"`
	Error string `json:"error" example:"Incorrect email"`
}

type HttpForbidden struct {
	Code  int    `json:"code" example:"403"`
	Error string `json:"error" example:"Your account is locked / Your account is expired"`
}

type NotFoundError struct {
	Code  int    `json:"code" example:"404"`
	Error string `json:"error" example:"User not found"`
}

type ServerError struct {
	Code  int    `json:"code" example:"500"`
	Error string `json:"error" example:"Error generating token"`
}
