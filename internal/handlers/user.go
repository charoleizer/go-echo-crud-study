package handlers

import (
	"net/http"

	"github.com/charoleizer/go-echo-crud-study/internal/cache"
	"github.com/charoleizer/go-echo-crud-study/internal/models"

	"github.com/labstack/echo/v4"
)

// Struct to supply interface
type user struct {
	*models.User
}

// Acessible by routes methods
type IUser interface {
	CreateUser(c echo.Context, u *models.User) error
}

// Interface object to expose methods
var User IUser = &user{&models.User{}}

// Create user and return status
func (h *user) CreateUser(c echo.Context, u *models.User) error {
	err := c.Bind(u)
	if err != nil {
		return err
	}

	cache.Users[u.ID] = u
	cache.Seq++
	return c.JSON(http.StatusCreated, u)
}
