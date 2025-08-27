package main

import (
	"estudo-go/api"
	"estudo-go/internal/adapters/controllers"
	"estudo-go/internal/core/services"
	"estudo-go/internal/infrastructure/database"
	"estudo-go/pkg/logging"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := logging.NewLogger()

	db, err := database.NewConnection()
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	dbType := os.Getenv("DB_TYPE")

	userRepo, err := database.FactoryNewUserRepository(dbType, db, logger)
	if err != nil {
		logger.Error("Failed to initialize MySQL repository")
	}
	defer userRepo.Close()

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService, logger)

	e := echo.New()

	api.RegisterUserRoutes(e, userController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Server starting on port %s", port)
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
		logger.Error("Failed to start server: %v", err)
	}
}
