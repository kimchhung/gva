package dto

import (
	"backend/app/common/model"
)

// Requests & responses Data Transfer Object
type LoginResponse struct {
	Token string       `json:"token"`
	Admin *model.Admin `json:"admin"`
}

type RegisterResponse struct {
	Token string       `json:"token"`
	Admin *model.Admin `json:"admin"`
}
