package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mattn/echo-livereload"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year  uint64 `json:"year" validate:"required"`
		Month uint64 `json:"month" validate:"required"`
		Day   uint64 `json:"day" validate:"required"`
	}

	// PathQueryValidator ensures date inputs are valid.
	PathQueryValidator struct {
		validator *validator.Validate
	}

	Message struct {
		Path int `json:"path"`
	}
)

func main() {
	// Instantiate Echo.
	e := echo.New()

	// Set up middlewares.
	e.Use(middleware.RequestID()) // https://echo.labstack.com/middleware/request-id
	e.Use(middleware.Logger())    // https://echo.labstack.com/middleware/logger
	e.Use(middleware.Recover())   // https://echo.labstack.com/middleware/recover
	e.Use(livereload.LiveReload())

	// Configure validation for PathQuery.
	// https://echo.labstack.com/guide/request#validate-data
	e.Validator = &PathQueryValidator{validator: validator.New()}

	// POST /path
	e.POST("/path", func(c echo.Context) (err error) {
		pq := &PathQuery{}
		if err = c.Bind(pq); err != nil {
			return
		}
		if err = c.Validate(pq); err != nil {
			return
		}

		final := CalculateLifePath(pq.Day, pq.Month, pq.Year)
		return c.JSON(http.StatusOK, final)
	})

	e.Logger.Fatal(e.Start(":1234"))
}

// Validate ensures data sent to the server is valid,
// and informs the user of invalide submissions.
func (cv *PathQueryValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Sum adds each digit of a number.
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

func CalculateLifePath(d uint64, m uint64, y uint64) int {
	// Step 1: Sum each individual date portion.
	day := sum(d)
	month := sum(m)
	year := sum(y)

	// Step 2: Sum again if a date portion is greater than 9 and does not equal 11.
	day = total(day)
	month = total(month)
	year = total(year)

	// Step 3: Add and return
	return day + month + year
}
