package apikey

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cappit/internal/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c echo.Context) error {
	keys, err := h.service.ListAPIKeys(c.Request().Context())
	if err != nil {
		logger.Error("failed to list API keys", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "failed to list API keys")
	}
	return c.JSON(http.StatusOK, keys)
}

func (h *Handler) POST(c echo.Context) error {
	var req CreateAPIKeyRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("invalid request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	apiKey, err := h.service.CreateAPIKey(c.Request().Context(), req)
	if err != nil {
		logger.Error("failed to create API key", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "failed to create API key")
	}
	return c.JSON(http.StatusCreated, apiKey)
}

func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")

	err := h.service.DeleteAPIKey(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "API key not found")
		}

		logger.Error("DB error while deleting API key", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "failed to delete API key")
	}

	return c.NoContent(http.StatusNoContent)
}
