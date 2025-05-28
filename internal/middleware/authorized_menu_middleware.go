package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"mkk_notification_bot/internal/config"
	"mkk_notification_bot/internal/handlers"
	"mkk_notification_bot/internal/models"
	"sync"
)

var initOnce = &sync.Once{}

func LoadCommand() {
	initOnce.Do(func() {
		config.Init()
	})
}

func authorizedMenuMiddleware(bot *tgbotapi.BotAPI, chatID int64, text string, user *models.UserState, logger *slog.Logger) {
	LoadCommand()

	if text == "" {
		text = "/start"
	}

	cmd, exists := config.Commands[text]
	if !exists {
		handlers.SendResponse(bot, chatID, "Неизвестная команда. Используйте /help", nil, logger)
	}

	handlers.SendResponse(bot, chatID, cmd.Text, cmd.Keyboard, logger)
}
