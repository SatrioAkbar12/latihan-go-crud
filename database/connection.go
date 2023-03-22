package database

import (
	"fmt"

	"github.com/SatrioAkbar12/latihan-go-crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	cred := "root@tcp(localhost:3306)/go-crud?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(cred), &gorm.Config{})

	if err != nil {
		fmt.Println("Conection error to database")
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
