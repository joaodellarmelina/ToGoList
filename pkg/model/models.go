// Contém as Structs que representam as tabelas no BD
package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Titulo     string `json:"titulo"`
	Descricao  string `json:"descricao"`
	Vencimento string `json:"vencimento"`
	Status     bool   `json:"status" gorm:"default: false"`
}
