package handler

import (
	"net/http"

	"github.com/Chatbot/logger"
	"github.com/Chatbot/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Message struct {
	Input string `json:"message"`
}

type Response struct {
	Reply string `json:"response"`
}

type ChatbotHandler struct {
	botService service.BotService
}

func NewChatbotHandler(botService service.BotService) *ChatbotHandler {
	return &ChatbotHandler{botService: botService}
}

func (h *ChatbotHandler) HandleChat(c echo.Context) (*Response, error) {
	msg := new(Message)
	if err := c.Bind(msg); err != nil {
		logger.Error("Failed to bind message", zap.Error(err))
		return nil, NewErrorResponse("invalid input", http.StatusBadRequest)
	}

	if msg.Input == "" {
		logger.Error("Input is empty")
		return nil, NewErrorResponse("input is empty", http.StatusBadRequest)
	}

	resp := h.botService.GetReply(msg.Input)
	logger.Info("Response generated", zap.String("input", msg.Input), zap.String("response", resp))
	return &Response{Reply: resp}, nil
}
