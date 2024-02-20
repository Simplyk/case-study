package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() *echo.Echo {
	e := echo.New()
	apiGroup := e.Group("/api/v1")
	ApiV1Controller(apiGroup)
	return e
}

func TestApiV1Controller(t *testing.T) {
	t.Parallel()
	e := setup()

	testCases := []struct {
		name   string
		method string
		path   string
		code   int
	}{
		{
			name:   "GET /api/v1",
			method: "GET",
			path:   "/api/v1",
			code:   http.StatusOK,
		},
		{
			name:   "GET /",
			method: "GET",
			path:   "/",
			code:   http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.code, rec.Code)
		})
	}
}
