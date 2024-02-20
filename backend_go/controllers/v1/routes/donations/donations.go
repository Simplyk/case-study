package donations

import (
	"net/http"
	"zbackend/data/query"

	"github.com/labstack/echo/v4"
)

type PaginationQueryParams struct {
	Limit  int `query:"limit" validate:"min=0,max=100"`
	Offset int `query:"offset" validate:"min=0,max=100"`
}

// GetAll returns all donations or a paginated list of donations.
func GetAll(c echo.Context) error {
	var queryParams PaginationQueryParams
	if err := c.Bind(&queryParams); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&queryParams); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// paginated query
	if queryParams.Limit != 0 {
		donations := query.GetDonationsPaginated(queryParams.Limit, queryParams.Offset)
		return c.JSON(http.StatusOK, donations)
	}

	donations, err := query.GetDonations()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, donations)
}
