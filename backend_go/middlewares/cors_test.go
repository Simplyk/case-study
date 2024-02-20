package middlewares

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCors(t *testing.T) {
	t.Run("should return a function", func(t *testing.T) {
		got := Cors(echo.HandlerFunc(func(c echo.Context) error {
			return nil
		}))
		if got == nil {
			t.Errorf("Cors() = %v, want %v", got, nil)
		}
	})
	t.Run("should call next", func(t *testing.T) {
		var called bool
		next := echo.HandlerFunc(func(c echo.Context) error {
			called = true
			return nil
		})
		Cors(next)(echo.New().NewContext(nil, nil))
		if !called {
			t.Errorf("Cors() should call next")
		}
	})
}
