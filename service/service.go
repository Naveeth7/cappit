package service

import "strings"

type botService struct{}

func NewBotService() BotService {
	return &botService{}
}

func (b *botService) GetReply(input string) string {
	lower := strings.ToLower(input)

	switch {
	case strings.Contains(lower, "hello"):
		return "Hi there! How can I help you today?"
	case strings.Contains(lower, "name"):
		return "I'm GoBot, your AI assistant built with Go!"
	case strings.Contains(lower, "bye"):
		return "Goodbye! Have a great day!"
	default:
		return "I'm not sure how to respond to that. Can you rephrase?"
	}
}
