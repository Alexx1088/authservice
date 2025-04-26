package dto

type RegisterRequestDTO struct {
	Name     string `json:"name" validate:"required, min=2, max=50"`
	Surname  string `json:"surname" validate:"required,min=2, max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
