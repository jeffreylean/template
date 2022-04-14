package router

import (
	"github.com/jeffreylean/template/echo-web-server/api/internal/bootstrap"
	"github.com/jeffreylean/template/echo-web-server/api/internal/handler"
	"github.com/labstack/echo/v4"
)

func V1(e *echo.Echo, bs *bootstrap.Bootstrap) {
	h := handler.New(bs)
	v1 := e.Group("/v1")
	v1.GET("/health", h.APIHealthCheck)
}
