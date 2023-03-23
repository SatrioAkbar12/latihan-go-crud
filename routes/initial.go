package routes

import (
	"github.com/SatrioAkbar12/latihan-go-crud/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteInitial() {
	e := echo.New()
	e.Validator = &models.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	UserRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
