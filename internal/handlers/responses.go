package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

func SendResponse(bot *tgbotapi.BotAPI, chatID int64, text string, keyboard *tgbotapi.ReplyKeyboardMarkup, logger *slog.Logger) {
	msg := tgbotapi.NewMessage(chatID, text)

	if keyboard != nil {
		msg.ReplyMarkup = keyboard
	}

	if _, err := bot.Send(msg); err != nil {
		logger.Error("error sending message", slog.String("err", err.Error()))
	}
}
