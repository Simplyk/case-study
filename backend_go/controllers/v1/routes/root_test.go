package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() *echo.Echo {
	e := echo.New()
	return e
}

func TestRoot(t *testing.T) {
	t.Parallel()
	e := setup()

	testCases := []struct {
		name     string
		method   string
		path     string
		expected string
	}{
		{
			name:     "GET /",
			method:   http.MethodGet,
			path:     "/",
			expected: "Hello, World!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e.GET("/", Root)

			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, tc.expected, rec.Body.String())
		})
	}
}
