package server

import (
	"net/http"

	bs "github.com/jeffreylean/template/echo-web-server/api/internal/bootstrap"
	"github.com/jeffreylean/template/echo-web-server/api/internal/router"
	"github.com/labstack/echo/v4"
)

func Start(port string) {
	bs := bs.New()
	e := echo.New()
	e.Validator = bs.Validator

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "info", "API WORKING FINE.")
	})

	router.V1(e, bs)

	e.Logger.Fatal(e.Start(":" + port))
}
