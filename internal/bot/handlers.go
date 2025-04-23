package bot

import (
	"fmt"
	"strconv"

	"github.com/bambutcha/bmi-calculator/internal/bmi"
	"github.com/bambutcha/bmi-calculator/internal/fsm"
	"github.com/bambutcha/bmi-calculator/pkg/keyboard"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		if message.Command() == "start" {
			b.handleStart(message)
		}
		return
	}

	state := b.storage.GetState(message.Chat.ID)
	if state == nil {
		return
	}

	switch state.State {
	case fsm.WaitingForHeight:
		b.handleHeight(message)
	case fsm.WaitingForWeight:
		b.handleWeight(message)
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID,
		"👋 Привет! Я бот для расчета индекса массы тела (ИМТ).\n\n"+
			"Пожалуйста, введите ваш рост в сантиметрах (например, 175):")

	b.storage.SetState(message.Chat.ID, fsm.NewUserState())
	b.API.Send(msg)
}

func (b *Bot) handleHeight(message *tgbotapi.Message) {
	height, err := strconv.ParseFloat(message.Text, 64)
	if err != nil || height <= 0 || height > 250 {
		msg := tgbotapi.NewMessage(message.Chat.ID,
			"Пожалуйста, введите корректное значение роста в сантиметрах (от 1 до 250):")
		b.API.Send(msg)
		return
	}

	state := b.storage.GetState(message.Chat.ID)
	state.Height = height
	state.State = fsm.WaitingForWeight
	b.storage.SetState(message.Chat.ID, state)

	msg := tgbotapi.NewMessage(message.Chat.ID,
		"Теперь введите ваш вес в килограммах (например, 70):")
	b.API.Send(msg)
}

func (b *Bot) handleWeight(message *tgbotapi.Message) {
	weight, err := strconv.ParseFloat(message.Text, 64)
	if err != nil || weight <= 0 || weight > 500 {
		msg := tgbotapi.NewMessage(message.Chat.ID,
			"Пожалуйста, введите корректное значение веса в килограммах (от 1 до 500):")
		b.API.Send(msg)
		return
	}

	state := b.storage.GetState(message.Chat.ID)
	bmi, category := bmi.CalculateBMI(state.Height, weight)

	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf("Ваш ИМТ: %.2f\n\nКатегория: %s\n\nХотите рассчитать ИМТ снова?",
			bmi, category))
	msg.ReplyMarkup = keyboard.GetRestartInlineKeyboard()
	b.API.Send(msg)

	b.storage.ClearState(message.Chat.ID)
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "restart":
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
			"Пожалуйста, введите ваш рост в сантиметрах (например, 175):")
		b.storage.SetState(callback.Message.Chat.ID, fsm.NewUserState())
		b.API.Send(msg)
	case "end":
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
			"Спасибо за использование бота! Если захотите рассчитать ИМТ снова, отправьте команду /start")
		b.API.Send(msg)
	}
	b.API.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, ""))
}
