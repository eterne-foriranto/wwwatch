package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func getBot() *tgbotapi.BotAPI {
	token := getConfigValue("telegram", "token")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true
	return bot
}

func send(text string) {
	chatID, err := strconv.Atoi(getConfigValue("telegram", "chat_id"))

	if err != nil {
		panic(err)
	}

	message := tgbotapi.NewMessage(int64(chatID), text)
	bot := getBot()
	bot.Send(message)
}
