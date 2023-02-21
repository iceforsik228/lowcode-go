package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5905 user=postgres dbname=railway password=pC4MDh6BP7OFOFpAcjyT sslmode=disable")

	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Snippet{})

	return db
}
