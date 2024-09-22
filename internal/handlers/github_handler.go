package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/sample", sampleFunc)
}

func sampleFunc(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample api")
}
