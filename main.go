package main

import (
	"github.com/sarang-pratham/visual-flow-backend/internal/config"
	"github.com/sarang-pratham/visual-flow-backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.LoadConfig()
	handlers.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(cfg.ServerPort))
}
