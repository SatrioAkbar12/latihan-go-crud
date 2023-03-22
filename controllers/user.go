package controllers

import (
	"fmt"
	"net/http"

	"github.com/SatrioAkbar12/latihan-go-crud/database"
	"github.com/SatrioAkbar12/latihan-go-crud/models"
	"github.com/labstack/echo"
)

func GetUsers(e echo.Context) error {
	db := database.DbConnection()

	users := []models.User{}

	db.Find(&users)

	return e.JSON(http.StatusOK, users)
}

func GetUserByID(e echo.Context) error {
	db := database.DbConnection()

	user := models.User{}
	id := e.Param("id")

	db.First(&user, id)

	return e.JSON(http.StatusOK, user)
}

func CreateUser(e echo.Context) error {
	db := database.DbConnection()

	user := models.User{}

	e.Bind(&user)
	db.Create(&user)

	fmt.Println(user)

	return e.JSON(http.StatusOK, user)
}

func UpdateUser(e echo.Context) error {
	db := database.DbConnection()

	user := models.User{}
	id := e.Param("id")

	db.First(&user, id)
	e.Bind(&user)
	db.Save(&user)

	fmt.Println(user)

	return e.JSON(http.StatusOK, user)
}

func DeleteUser(e echo.Context) error {
	db := database.DbConnection()

	user := models.User{}
	id := e.Param("id")

	db.First(&user, id)
	db.Delete(&user)

	return e.JSON(http.StatusOK, user)
}
