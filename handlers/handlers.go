package handlers

import (
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/KitlerUA/NNCompanionBot/matching"
	"github.com/KitlerUA/NNCompanionBot/utils"
	"github.com/yanzay/tbot"
)

func ComputeAnswer(message *tbot.Message) {
	defer func() {
		if r := recover(); r != nil {
			if !utils.OnlyCyrylik(message.Text()) {
				message.Reply("I`m so stupid ;( . I dont'n understand you")
			}
			loger.Log.Warning("Recovered after message: ", message)
		}
	}()
	rep := matching.QAGeneral[matching.CMGeneral.Closest(message.Text())]
	if rep != "" {
		message.Reply(rep)
	} else {
		message.Reply("I have nothing to say")
	}

	loger.Log.Info("To ", message.From.UserName, " ", message.From.FirstName, " ", message.From.LastName, " : ", rep)

}
