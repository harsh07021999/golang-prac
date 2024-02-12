package dtos

type UserRequest struct {
	UserName string `json:"user_name"`
}

type UserResponse struct {
	Email string `json:"email"`
}
