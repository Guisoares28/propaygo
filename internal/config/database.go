package config

import (
	"log"
	"propay/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=admin password=123 dbname=db_propay port=8080 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}

	database.AutoMigrate(&model.User{}, &model.Account{})

	DB = database
}
