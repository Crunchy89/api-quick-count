package responses

type RoleResponse struct {
	Base
	Name   string `json:"name"`
	Edit   bool   `json:"edit"`
	Create bool   `json:"create"`
	Delete bool   `json:"delete"`
}
