package tgclient

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// config
const UPDATE_TIMEOUT = 30

// commands
const HELP = "help"
const GREET = "greet"

func Run() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Panic(err)
	}
	respondToUpdates(bot)
}

func respondToUpdates(bot *tgbotapi.BotAPI) {
	updateConfig := getUpdateConfig()
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message.IsCommand() {
			respondToCommands(bot, update)
		}
	}
}

func getUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UPDATE_TIMEOUT
	return updateConfig
}

func respondToCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = createCommandResponse(update.Message.Command())

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func createCommandResponse(command string) string {
	switch command {
	case HELP:
		return fmt.Sprintf("Available commands are:\n /%s\n /%s", HELP, GREET)
	case GREET:
		return "Greetings"
	default:
		return fmt.Sprintf("I don't know that command, use /%s for list commands", HELP)
	}
}