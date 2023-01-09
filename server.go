package main

import (
	"net/http"
	"smart-home-backend/route"

	"github.com/labstack/echo/v4"

	"smart-home-backend/database"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	e := echo.New()

	// Initialize Database
	database.Connect()
	database.Migrate()

	route.Routing(e.Group("/api"))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
