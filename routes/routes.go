package routes

import (
	"net/http"

	"github.com/cappit/auth"
	"github.com/cappit/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, chatHandler *handler.ChatbotHandler) {
	e.POST("/login", func(c echo.Context) error {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		// Mock Auth AD logic
		if req.Username == "admin" && req.Password == "admin" {
			token, err := auth.GenerateToken(req.Username)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, echo.Map{
				"token": token,
			})
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	})

	api := e.Group("/api")
	api.Use(auth.JWTMiddleware)

	api.POST("/chat", func(c echo.Context) error {
		resp, err := chatHandler.HandleChat(c)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})
}
