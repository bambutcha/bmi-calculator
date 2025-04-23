package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bambutcha/bmi-calculator/internal/bot"
)

func main() {
	b, err := bot.NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", b.API.Self.UserName)

	// Обработка сигналов для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := b.Start(); err != nil {
			log.Printf("Error running bot: %v", err)
		}
	}()

	<-sigChan
	log.Println("Shutting down...")
}
