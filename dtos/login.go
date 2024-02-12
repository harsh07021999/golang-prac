package dtos

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
