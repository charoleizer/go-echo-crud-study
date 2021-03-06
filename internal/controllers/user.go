package controllers

import (
	"github.com/charoleizer/go-echo-crud-study/internal/cache"
	"github.com/charoleizer/go-echo-crud-study/internal/handlers"
	"github.com/charoleizer/go-echo-crud-study/internal/models"

	"github.com/labstack/echo/v4"
)

// Invoke CreateUser interface. Cannot invoke directly by router
func CreateUser(c echo.Context) error {
	return handlers.User.CreateUser(c, &models.User{ID: cache.Seq})
}
