package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mkk_notification_bot/internal/models"
)

func authorizedMenuMiddleware(bot *tgbotapi.BotAPI, chatID int64, text string, user *models.UserState) {
	keyboard := getAuthorizedMenu()
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = false

	if text == "/menu" || text == "/start" || text == "" {
		msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
		return
	}

	switch text {
	case "📊 Информация о займе", "/loan":
		msg := tgbotapi.NewMessage(chatID, "Информация о вашем займе... (в разработке)")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case "❓ FAQ", "/faq":
		msg := tgbotapi.NewMessage(chatID, "FAQ: Как получить займ? Ответ: через личный кабинет.")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case "📞 Связь с оператором", "/operator":
		msg := tgbotapi.NewMessage(chatID, "Вы переведены к оператору. Ожидайте ответа...")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case "/help":
		msg := tgbotapi.NewMessage(chatID, "Доступные команды:\n/menu — главное меню\n/loan — информация о займе\n/faq — часто задаваемые вопросы\n/operator — связь с поддержкой")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Используйте доступные опции ниже.")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	}
}

func getAuthorizedMenu() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📊 Информация о займе"),
			tgbotapi.NewKeyboardButton("❓ FAQ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📞 Связь с оператором"),
		),
	)
}
