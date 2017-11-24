package main

import (

	"gopkg.in/telegram-bot-api.v4"
	"github.com/KitlerUA/NNCompanionBot/config"
	"github.com/KitlerUA/NNCompanionBot/loger"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig().Token)
	if err != nil {

		loger.Log.Panic(err)
	}

	bot.Debug = true

	loger.Log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		loger.Log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
