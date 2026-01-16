package dtos

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}
