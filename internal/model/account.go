package model

import "gorm.io/gorm"

/*
Não é necessário definir uma ForeignKey o gorm já entende, por isso
passamos o ID do usuário somente como um campo padrão
*/
type Account struct {
	gorm.Model
	UserID uint    `gorm:"not null"`
	Title  string  `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	Status string  `gorm:"not null;default:'pending'"`
}
