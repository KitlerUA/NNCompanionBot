package handlers

import "github.com/KitlerUA/NNCompanionBot/matching"

func ComputeAnswer(message string) string {
	k := matching.CMForCategories[matching.CM.Closest(message)].Closest(message)
	for i := range matching.Categories {
		for j := range matching.Categories[i].QAs {
			if k == j {
				return matching.Categories[i].QAs[j]
			}
		}
	}
	return "I have nothing to say"
}
