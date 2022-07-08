package routes

import (
	"github.com/charoleizer/go-echo-crud-study/internal/controllers"

	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)

}
