package responses

type UserResponse struct {
	Username string `json:"username"`
	RoleUUID string `json:"role_uuid"`
	IsOnline bool   `json:"is_online"`
	Active   bool   `json:"active"`
}

type UserResponseJWT struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	IsOnline bool   `json:"is_online"`
	Active   bool   `json:"active"`
}
