package dto

type AccountRequestDto struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
