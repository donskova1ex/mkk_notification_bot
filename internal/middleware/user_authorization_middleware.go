package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"mkk_notification_bot/internal"
	"mkk_notification_bot/internal/models"
	"mkk_notification_bot/internal/utils"
	"regexp"
	"strings"
)

func userPhoneNumberAuthorization(
	bot *tgbotapi.BotAPI,
	chatID int64,
	msg *tgbotapi.Message,
	state *models.UserState,
	logger *slog.Logger) {

	text := strings.TrimSpace(msg.Text)
	if text == "" {
		_, err := bot.Send(tgbotapi.NewMessage(chatID, internal.NumberInputMessage))
		if err != nil {
			logger.Error(
				"error sending authorization message ",
				slog.String("err", err.Error()),
			)
			return
		}
	}
	normalizedNumber, ok := utils.NormalizePhone(text, logger)
	if !ok {
		_, err := bot.Send(tgbotapi.NewMessage(chatID, internal.WrongNumberFormat))
		if err != nil {
			logger.Error(
				"error sending number format message",
				slog.String("err", err.Error()),
				)
		}
	}

	state.Phone = text

	ok, err :=

}
