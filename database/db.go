package database

import (
	"log"

	"github.com/pedropaccola/go-apigin-alura/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error: can't connect to the database")
	}
	DB.AutoMigrate(&models.Student{})
}
