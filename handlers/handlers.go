package handlers

import (
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/KitlerUA/NNCompanionBot/matching"
	"github.com/yanzay/tbot"
)

func ComputeAnswer(message *tbot.Message) {
	defer func() {
		if r := recover(); r != nil {
			loger.Log.Warning("Recovered after message: ", message)
		}
	}()
	k := matching.CMForCategories[matching.CM.Closest(message.Vars["text"])].Closest(message.Vars["text"])
	for i := range matching.Categories {
		for j := range matching.Categories[i].QAs {
			if k == j {
				message.Reply(matching.Categories[i].QAs[j])
				loger.Log.Info("To ", message.From.UserName, " ", message.From.FirstName, " ", message.From.LastName, " : ", matching.Categories[i].QAs[j])
				break
			}
		}
	}
}
