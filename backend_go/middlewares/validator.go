package middlewares

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validator for echo
type Validator struct {
	validator *validator.Validate
}

// Validation middleware.
// Augmente le HandlerFunc echo avec un Validator dans le Context.
// Retourne le HandlerFunc avec l'ajout.
func Validation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Echo().Validator = &Validator{validator: validator.New()}
		return next(c)
	}
}

// Méthode à appeler pour valider les données prédéfinies dans les struct.
// Retourne une erreur si la validation échoue, nil sinon.
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		acc := err.(validator.ValidationErrors)

		return echo.NewHTTPError(http.StatusBadRequest, acc)
	}

	return nil
}
