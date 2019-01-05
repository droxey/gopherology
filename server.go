package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year int `json:"year" validate:"required"`
		Month int `json:"month" validate:"required"`
		Day int `json:"day" validate:"required"`
	}

	// PathQueryValidator ensures date inputs are valid.
	PathQueryValidator struct {
		validator *validator.Validate
	}
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &PathQueryValidator{validator: validator.New()}

	e.POST("/path", func(c echo.Context) (err error) {
		pq := new(PathQuery)
		if err = c.Bind(pq); err != nil {
			return
		}
		if err = c.Validate(pq); err != nil {
			return
		}
		return c.JSONPretty(http.StatusOK, pq, "  ")
	})
	e.Logger.Fatal(e.Start(":9999"))
}

// Validate ensures data sent to the server is valid,
// and informs the user of invalide submissions.
func (cv *PathQueryValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
