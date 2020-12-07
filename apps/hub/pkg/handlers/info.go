package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type InfoHandler struct {
}

func (h *InfoHandler) Routes(router *echo.Group) {
	router.GET("/health", h.Health)
}

// Health ...
func (h *InfoHandler) Health(context echo.Context) error {

	return context.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
