// Responsável por fazer a Conexão com o Banco de dados
package database

import (
	"github.com/RamsesMartins/ToGoList/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(model.Task{})
	return db
}
