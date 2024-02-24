package dto

// Requests & responses Data Transfer Object
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
