package main

import (
	"os"

	"github.com/Chatbot/handler"
	"github.com/Chatbot/logger"
	"github.com/Chatbot/middleware"
	"github.com/Chatbot/routes"
	"github.com/Chatbot/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.Init()
	e := echo.New()

	middleware.Logger(e)

	botService := service.NewBotService()
	chatHandler := handler.NewChatbotHandler(botService)

	routes.RegisterRoutes(e, chatHandler)
	if err := godotenv.Load("config/.env"); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Print("Starting GoBot on http://localhost:" + os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
