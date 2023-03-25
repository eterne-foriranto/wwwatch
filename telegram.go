package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

const TelegramTokenFileName = "tg_token.dat"
const ChatID = 195336731

func getTgToken(fileName string) string {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	b := make([]byte, 45)
	_, err = f.Read(b)

	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func getBot() *tgbotapi.BotAPI {
	token := getTgToken(TelegramTokenFileName)
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true
	return bot
}

func send(text string) {
	message := tgbotapi.NewMessage(ChatID, text)
	bot := getBot()
	bot.Send(message)
}
