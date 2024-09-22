package main

import (
	"github.com/sarang-pratham/visual-flow-backend/internal/config"
	"github.com/sarang-pratham/visual-flow-backend/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Visual Flow APIs
// @version 1.0
// @description APIs for all the calculations/features.
// @termsOfService http://swagger.io/terms/
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	api := e.Group("visual-flow-service/api/v1")
	cfg := config.LoadConfig()

	handlers.RegisterRoutes(api)
	e.Logger.Fatal(e.Start(cfg.ServerPort))
}
