package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := runSever(); err != nil {
		// ライブラリ使いたい
		log.Fatal(err)
	}
}

func runSever() error {
	e := echo.New()
	e = setMiddleware(e)
	e = setRoutes(e)

	return e.Start(address())
}

func setMiddleware(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

func setRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "pong"})
	})
	return e
}

func address() string {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	return ":" + httpPort
}
