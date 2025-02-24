package dto

import "github.com/gva/internal/ent"

// Requests & responses Data Transfer Object
type LoginResponse struct {
	Token string     `json:"token"`
	Admin *ent.Admin `json:"admin"`
}

type RegisterResponse struct {
	Token string     `json:"token"`
	Admin *ent.Admin `json:"admin"`
}
