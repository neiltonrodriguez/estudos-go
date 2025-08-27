package api

import (
	"estudo-go/internal/adapters/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, userController *controllers.UserController) {
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	usersGroup := e.Group("/users")

	usersGroup.POST("", userController.CreateUser)
	usersGroup.GET("", userController.GetUserByEmail)
}
