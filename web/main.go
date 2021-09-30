package main

import (
	"fmt"

	"github.com/coil398/go-web/web/Openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	server := &server{}
	Openapi.RegisterHandlersWithBaseURL(e, server, "v1")
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:3000")))
}
