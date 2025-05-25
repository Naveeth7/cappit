package main

import (
	"net/http"

	"github.com/Chatbot/handler"
	"github.com/Chatbot/logger"
	"github.com/Chatbot/middleware"
	"github.com/Chatbot/service"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.Init()
	e := echo.New()

	middleware.Logger(e)

	botService := service.NewBotService()
	chatHandler := handler.NewChatbotHandler(botService)

	e.POST("/chat", func(c echo.Context) error {
		resp, err := chatHandler.HandleChat(c)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Print("Starting GoBot on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
