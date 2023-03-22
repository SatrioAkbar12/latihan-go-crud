package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteInitial() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	UserRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
