package main

import (
	"github.com/KitlerUA/NNCompanionBot/config"
	"github.com/KitlerUA/NNCompanionBot/handlers"
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/KitlerUA/NNCompanionBot/matching"
	"github.com/yanzay/tbot"
)

func main() {

	bot, err := tbot.NewServer(config.Get().Token)
	if err != nil {
		loger.Log.Fatalf("Cannot get token from config file: %v", err)
	}
	matching.PrepareDataForMatching()
	bot.AddMiddleware(logMid)
	bot.HandleFunc("{text}", handlers.ComputeAnswer)
	err = bot.ListenAndServe()
	loger.Log.Fatal(err)
}

func logMid(f tbot.HandlerFunction) tbot.HandlerFunction {
	return func(m *tbot.Message) {
		loger.Log.Info("From ", m.From.UserName, " ", m.From.FirstName, " ", m.From.LastName, " : ", m.Text())
		f(m)
	}
}
