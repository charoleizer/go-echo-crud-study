package main

import (
	"github.com/charoleizer/go-echo-crud-study/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.MapRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
