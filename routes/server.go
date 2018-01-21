package routes

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// StartServer is registering and starting the web server
func StartServer() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://pierrez.github.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/api/v0/assets", getAssets)
	e.POST("/api/v0/assets/:asset", buyAsset)
	e.POST("/api/v0/money", addMoney)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
