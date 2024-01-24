package responses

type AuthResponse struct {
	Token string           `json:"token"`
	User  *UserResponseJWT `json:"user"`
}
