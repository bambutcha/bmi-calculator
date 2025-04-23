package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bambutcha/bmi-calculator/internal/fsm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

type Bot struct {
	API     *tgbotapi.BotAPI
	storage *fsm.StateStorage
}

func NewBot() (*Bot, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	api, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &Bot{
		API:     api,
		storage: fsm.NewStateStorage(),
	}, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.API.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("failed to get updates channel: %w", err)
	}

	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			b.handleCallback(update.CallbackQuery)
		}
	}

	return nil
}
