package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"mkk_notification_bot/internal/models"
)

var userStates = make(map[int64]*models.UserState)

func MessageMiddleware(bot *tgbotapi.BotAPI, update tgbotapi.Update,  logger *slog.Logger) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID

	state, ok := userStates[chatID]
	if !ok {
		state = &models.UserState{
			UserID: userID,
			State:  "waiting_authorization",
		}
		userStates[chatID] = state
	}

	slog.Info("Message received from [userID: %d]", userID)

	switch state.State {
	case "waiting_authorization":

	case :
		

	}

}
