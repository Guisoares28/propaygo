package dto

type UserRequestDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponseDto struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
