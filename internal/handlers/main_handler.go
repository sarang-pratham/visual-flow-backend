package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/sarang-pratham/visual-flow-backend/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Group) {
	e.GET("/", healthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/sample", sampleFunc)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags HealthCheck
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

func sampleFunc(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample api")
}
