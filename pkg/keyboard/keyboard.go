package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetRestartInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Рассчитать ИМТ снова", "restart"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завершить", "end"),
		),
	)
}
