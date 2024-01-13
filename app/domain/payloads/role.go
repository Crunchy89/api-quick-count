package payloads

import "github.com/google/uuid"

type CreateRolePayload struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Edit   bool   `json:"edit" form:"edit"`
	Create bool   `json:"create" form:"create"`
	Delete bool   `json:"delete" form:"delete"`
}

type UpdateRolePayload struct {
	Name   string `json:"name" form:"name"`
	Edit   bool   `json:"edit" form:"edit"`
	Create bool   `json:"create" form:"create"`
	Delete bool   `json:"delete" form:"delete"`
}

type DeleteRolePayload struct {
	UUID uuid.UUID `json:"uuid" form:"uuid" validate:"required"`
}
