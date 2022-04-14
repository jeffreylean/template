package handler

import (
	"net/http"

	"github.com/jeffreylean/template/echo-web-server/api/internal/bootstrap"
	"github.com/jeffreylean/template/echo-web-server/api/internal/repository"
	"github.com/labstack/echo/v4"
)

// Handler: Handler consists of all request handler.
type Handler struct {
	repository *repository.Repository
}

func New(bs *bootstrap.Bootstrap) *Handler {
	handler := &Handler{
		repository: bs.Repository,
	}
	return handler
}

// APIHealthCheck : Health check handler for api endpoint.
func (h Handler) APIHealthCheck(c echo.Context) error {
	return c.Render(http.StatusOK, "info", "API IS FINE")
}
