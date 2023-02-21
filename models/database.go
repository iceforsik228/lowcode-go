package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=lowcode password="+os.Getenv("PGPASSWORD")+" sslmode=disable")

	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Snippet{})

	return db
}
