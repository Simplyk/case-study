package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors middleware.
// Accept CORS pour localhost:3000.
// Retourne le HandlerFunc avec l'ajout.
func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Echo().Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
		}))
		return next(c)
	}
}
