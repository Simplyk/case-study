package main

import (
	"zbackend/controllers/v1"
	"zbackend/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ZBackend API.
// This is a sample server for ZBackend API.
// @version v1
// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middlewares.Cors)
	e.Use(middlewares.Validation)

	apiGroup := e.Group("/api/v1")
	controllers.ApiV1Controller(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
