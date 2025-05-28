package middleware

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"mkk_notification_bot/internal"
	"mkk_notification_bot/internal/models"
	"mkk_notification_bot/services"
	"sync"
)

var (
	userStates = make(map[int64]*models.UserState)
	mu         sync.Mutex
)

func MessageMiddleware(
	ctx context.Context,
	cds *services.ClientDataService,
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *slog.Logger,
) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID

	// Защита доступа к userStates
	mu.Lock()
	state, ok := userStates[chatID]
	if !ok {
		state = &models.UserState{
			UserID: userID,
			State:  internal.WaitingAuthorization,
		}
		userStates[chatID] = state
	}
	mu.Unlock()

	logger.Info("Message received", "userID", userID, "chatID", chatID)

	if update.Message.IsCommand() && update.Message.Command() == "start" {
		msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Пожалуйста, введите свой номер телефона и нажмите «Отправить номер телефона».")
		bot.Send(msg)
		return
	}
	logger.Info("Current user state", "chatID", chatID, "state", state.State)
	switch state.State {
	case internal.WaitingAuthorization:
		userPhoneNumberAuthorization(ctx, cds, bot, chatID, update.Message, state, logger)
	case internal.Authorized:
		authorizedMenuMiddleware(bot, chatID, update.Message.Text, state, logger)
	}
}
