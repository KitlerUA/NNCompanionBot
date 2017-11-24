package main

import (
	"time"

	"github.com/KitlerUA/NNCompanionBot/config"
	"github.com/KitlerUA/NNCompanionBot/handlers"
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/KitlerUA/NNCompanionBot/matching"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Get().Token)
	if err != nil {

		loger.Log.Panic(err)
	}
	loger.Log.Info("Preparing data")
	startTime := time.Now().Nanosecond()
	if err = matching.PrepareDataForMatching(); err != nil {
		loger.Log.Panicf("Cannot prepare matching data: %v", err)
	}
	loger.Log.Infof("Data prepared successfully. Took %v sec", float64(time.Now().Nanosecond()-startTime)/1000000000)

	bot.Debug = false

	loger.Log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		loger.Log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID
		var msg tgbotapi.MessageConfig
		if len(update.Message.Text) > 4 {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, handlers.ComputeAnswer(update.Message.Text))
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Please, type more then 3 characters")
		}
		loger.Log.Infof("[%s] %s", update.Message.From.UserName, msg.Text)
		bot.Send(msg)
	}
}
