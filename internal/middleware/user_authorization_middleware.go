package middleware

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"mkk_notification_bot/internal"
	"mkk_notification_bot/internal/models"
	"mkk_notification_bot/internal/utils"
	"mkk_notification_bot/services"
	"strings"
)

func userPhoneNumberAuthorization(
	ctx context.Context,
	cds *services.ClientDataService,
	bot *tgbotapi.BotAPI,
	chatID int64,
	msg *tgbotapi.Message,
	state *models.UserState,
	logger *slog.Logger,
) {

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
	if ok {
		ok, err := cds.ClientNumberFound(ctx, normalizedNumber)
		if err != nil {
			logger.Error(
				"failed to check if client number found",
				slog.String("err", err.Error()),
			)
			_, err := bot.Send(tgbotapi.NewMessage(chatID, internal.NumberCheckinError))
			if err != nil {
				slog.Error(
					"error sending number checkin answer",
					slog.String("err", err.Error()),
				)
			}
			return
		}
		if !ok {
			_, err := bot.Send(tgbotapi.NewMessage(chatID, internal.NumberNotFound))
			if err != nil {
				slog.Error(
					"error sending number not found",
					slog.String("err", err.Error()))
			}
		}

		if ok {
			state.Phone = normalizedNumber
			state.State = internal.Authorized
			_, err := bot.Send(tgbotapi.NewMessage(chatID, "Вы успешно авторизованы"))
			if err != nil {
				slog.Error("error sending message to user", slog.String("err", err.Error()))
			}
			authorizedMenuMiddleware(bot, chatID, "/menu", state, logger)
			return
		}
	}
}
