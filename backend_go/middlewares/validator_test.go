package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() *echo.Echo {
	e := echo.New()
	e.Use(Validation)
	e.GET("/", func(c echo.Context) error {
		type params struct {
			Param int `query:"param" validate:"min=1" json:"param"`
		}

		var p params
		if err := c.Bind(&p); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&p); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, p)
	})

	return e
}

func TestValidation(t *testing.T) {
	e := setup()

	testCases := []struct {
		name   string
		method string
		path   string
		code   int
		param  string
	}{
		{
			name:   "GET /?param=1",
			method: "GET",
			path:   "/?param=1",
			code:   http.StatusOK,
			param:  "{\"param\":1}\n",
		},
		{
			name:   "GET /?notparam=0 invalid",
			method: "GET",
			path:   "/?notparam=0",
			code:   http.StatusBadRequest,
			param:  "",
		},
		{
			name:   "GET /?param=0 invalid",
			method: "GET",
			path:   "/?param=0",
			code:   http.StatusBadRequest,
			param:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.code, rec.Code)
			if tc.code == http.StatusOK {
				assert.Equal(t, tc.param, rec.Body.String())
			}
		})
	}
}
