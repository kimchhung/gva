package dto

// Requests & responses Data Transfer Object
type LoginResponse struct {
	Token string `json:"token"`
	User  any    `json:"user"`
}
