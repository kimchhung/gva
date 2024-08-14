package dto

// Requests & responses Data Transfer Object
type LoginRequest struct {
	Username string `json:"username"  validate:"required,min=5,max=30"`
	Password string `json:"password"  validate:"required,min=6,max=30"`
}

type RegisterRequest struct {
	DisplayName string `json:"displayName"  validate:"required"`
	Username    string `json:"username"  validate:"required,min=5,max=30"`
	Password    string `json:"password"  validate:"required,min=6,max=30"`
}
