package main

import (
	"fmt"
	"net/http"

	"gopherology/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &models.PathQueryValidator{validator: models.Validator.New()}
	e.Use(middleware.RequestID()) // https://echo.labstack.com/middleware/request-id
	e.Use(middleware.Logger())    // https://echo.labstack.com/middleware/logger
	e.Use(middleware.Recover())   // https://echo.labstack.com/middleware/recover

	e.POST("/path", func(c echo.Context) (err error) {
		pq := &models.PathQuery{}
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

	e.Logger.Fatal(e.Start(":1234"))
}
