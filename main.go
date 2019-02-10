package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/droxey/gopherology/utils"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
			e.Logger.Fatal("$PORT must be set")
	}

	// Set up Echo, configure server side validation, and hook into middleware.
	e.Server.Addr = ":" + port
	e.Validator = &utils.PathQueryValidator{Validator: validator.New()}
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Attach a POST route to endpoint named /api/path.
	e.POST("/api/path", func(c echo.Context) (err error) {
		pq := &utils.PathQuery{}
		if err = c.Bind(pq); err != nil {
			return
		}
		if err = c.Validate(pq); err != nil {
			return
		}
		pq.Path = utils.CalculateLifePath(pq.Day, pq.Month, pq.Year)
		pq.Message = "Your Life Path Number is " + fmt.Sprint(pq.Path)
		return c.JSONPretty(http.StatusOK, pq, "  ")
	})

	// Gracefully shut down the server on interrupt.
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
