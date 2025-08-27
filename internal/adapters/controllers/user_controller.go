package controllers

import (
	"estudo-go/internal/core/domain"
	"estudo-go/internal/core/ports"
	"estudo-go/internal/core/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserController struct {
	userService services.UserService
	logger      ports.Logger
}

func NewUserController(userService services.UserService, logger ports.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	req := new(UserCreateRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := c.userService.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		if err == services.ErrUserAlreadyExists {
			return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}

		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	c.logger.Info("Success created user")

	user.Password = ""
	return ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUserByEmail(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Email query parameter is required"})
	}

	user, err := c.userService.GetUserByEmail(email)
	if err != nil {
		if err == domain.ErrUserNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	c.logger.Info("Success get user, email: ", email)

	user.Password = ""
	return ctx.JSON(http.StatusOK, user)
}
