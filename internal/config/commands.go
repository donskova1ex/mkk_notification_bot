package config

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Command struct {
	Text     string
	Keyboard *tgbotapi.ReplyKeyboardMarkup
}

var commands = make(map[string]Command)

func Init() {
	menuKeyboard := generateMenuKeyboard()
	commands = map[string]Command{
		"/menu": {
			Text:     "Выберите действие:",
			Keyboard: menuKeyboard,
		},
		"/start": {
			Text:     "Добро пожаловать! Выберите действие:",
			Keyboard: menuKeyboard,
		},
		"📊 Информация о займе": {
			Text:     "Информация о вашем займе... (в разработке)",
			Keyboard: menuKeyboard,
		},
		"/loan": {
			Text:     "Информация о вашем займе... (в разработке)",
			Keyboard: menuKeyboard,
		},
		"🧾 Активный займ": {
			Text:     "Активный займ: Такой-то номер",
			Keyboard: menuKeyboard,
		},
		"/active_loan": {
			Text:     "Активный займ: Такой-то номер",
			Keyboard: menuKeyboard,
		},
		"❓ FAQ": {
			Text:     "FAQ: Как получить займ? Ответ: через личный кабинет.",
			Keyboard: menuKeyboard,
		},
		"/faq": {
			Text:     "FAQ: Как получить займ? Ответ: через личный кабинет.",
			Keyboard: menuKeyboard,
		},
		"📞 Связь с оператором": {
			Text:     "Вы переведены к оператору. Ожидайте ответа...",
			Keyboard: menuKeyboard,
		},
		"/operator": {
			Text:     "Вы переведены к оператору. Ожидайте ответа...",
			Keyboard: menuKeyboard,
		},
		"/help": {
			Text:     "Доступные команды:\n/menu — главное меню\n/loan — информация о займе\n/faq — часто задаваемые вопросы\n/operator — связь с поддержкой",
			Keyboard: menuKeyboard,
		},
	}
}

func GetCommand(text string) (*Command, bool) {
	cmd, ok := commands[text]
	return &cmd, ok

}

func generateMenuKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📊 Информация о займе"),
			tgbotapi.NewKeyboardButton("🧾 Активный займ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("❓ FAQ"),
			tgbotapi.NewKeyboardButton("📞 Связь с оператором"),
		),
	)
	return &keyboard
}
