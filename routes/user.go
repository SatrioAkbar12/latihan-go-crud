package routes

import (
	"github.com/SatrioAkbar12/latihan-go-crud/controllers"
	"github.com/labstack/echo"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/user", controllers.GetUsers)
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:id", controllers.GetUserByID)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("user/:id", controllers.DeleteUser)
}
