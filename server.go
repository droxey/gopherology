package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year  uint64 `json:"year" validate:"required"`
		Month uint64 `json:"month" validate:"required"`
		Day   uint64 `json:"day" validate:"required"`
		Path  int    `json:"path"`
		Message string `json:"message"`
	}

	// PathQueryValidator ensures date inputs are valid.
	PathQueryValidator struct {
		validator *validator.Validate
	}
)

func main() {
	e := echo.New()
	e.Validator = &PathQueryValidator{validator: validator.New()}
	e.Use(middleware.RequestID()) // https://echo.labstack.com/middleware/request-id
	e.Use(middleware.Logger())    // https://echo.labstack.com/middleware/logger
	e.Use(middleware.Recover())   // https://echo.labstack.com/middleware/recover

	e.POST("/path", func(c echo.Context) (err error) {
		pq := &PathQuery{}
		if err = c.Bind(pq); err != nil {
			return
		}
		if err = c.Validate(pq); err != nil {
			return
		}

		pq.Path = CalculateLifePath(pq.Day, pq.Month, pq.Year)
		pq.Message = "Your Life Path Number is " + fmt.Sprint(pq.Path)
		return c.JSONPretty(http.StatusOK, pq, "  ")
	})

	e.Logger.Fatal(e.Start(":1234"))
}

// Validate ensures data sent to the server is valid,
// and informs the user of invalide submissions.
func (cv *PathQueryValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func sum(i uint64) (sum int) {
	b64 := uint64(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	return
}
func total(i int) (total int) {
	if i > 9 && i != 11 {
		total = sum(uint64(i))
	} else {
		total = i
	}
	return total
}

// CalculateLifePath: Sum each individual date portion.
// Sum again if a date portion is greater than 9 and does not equal 11.
// Return the totals.
func CalculateLifePath(d uint64, m uint64, y uint64) int {
	day := total(sum(d))
	month := total(sum(m))
	year := total(sum(y))
	return day + month + year
}
