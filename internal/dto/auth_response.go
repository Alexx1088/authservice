package dto

type AuthResponseDTO struct {
	Token   string `json:"token"`
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}
