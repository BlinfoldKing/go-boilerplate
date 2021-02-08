package auth

// RegisterRequest request for register
type RegisterRequest struct {
	Email            string  `json:"email" validate:"required,email"`
	Password         string  `json:"password" validate:"required"`
	CompanyContactID *string `json:"company_contact_id"`
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

// ResetPasswordSubmit request for reset request
type ResetPasswordSubmit struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
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
