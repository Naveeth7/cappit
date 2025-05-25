package service

// BotService defines the behavior of a bot.
type BotService interface {
	GetReply(input string) string
}
