package controllers

import (
	"zbackend/controllers/v1/routes"
	"zbackend/controllers/v1/routes/donations"

	"github.com/labstack/echo/v4"
)

// Enregistre les routes de l'API v1
func ApiV1Controller(g *echo.Group) {
	g.GET("", routes.Root)

	g.GET("/donations", donations.GetAll)
}
