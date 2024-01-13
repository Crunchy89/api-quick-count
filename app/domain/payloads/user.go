package payloads

import "github.com/google/uuid"

type CreateUserpayload struct {
	Username string    `json:"username" form:"username" validate:"required"`
	Password string    `json:"password" form:"password" validate:"required"`
	Email    string    `json:"email" form:"email" validate:"required"`
	RoleUUID uuid.UUID `json:"uuid" form:"uuid" validate:"required"`
}
