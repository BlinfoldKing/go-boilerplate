package auth

// RegisterRequest request for register
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginRequest request for register
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ResetPasswordRequest request for reset request
type ResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// VerifyActivationRequest requests for activation verification
type VerifyActivationRequest struct {
	Token string `json:"token" validate:"required,token"`
	Email string `json:"email" validate:"required,email"`
}

// ActivateAccountRequest request for account activation
type ActivateAccountRequest struct {
	Email string `json:"email" validate:"required,email"`
}
