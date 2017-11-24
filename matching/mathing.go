package matching

import (
	"encoding/json"
	"io/ioutil"

	"github.com/KitlerUA/NNCompanionBot/config"
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/schollz/closestmatch"
)

type Category struct {
	Name        string
	Description string
	QAs         map[string]string
}

var CM *closestmatch.ClosestMatch
var CMForCategories map[string]*closestmatch.ClosestMatch

var Categories []Category

func PrepareDataForMatching() error {
	data, err := ioutil.ReadFile(config.Get().MatchingDataFile)
	if err != nil {
		loger.Log.Warningf("Cannot find file with matching data: %v", err)
		return err
	}
	if err = json.Unmarshal(data, &Categories); err != nil {
		loger.Log.Warningf("Corrupted data in matching file: %v", err)
		return err
	}
	descriptions := make([]string, 0, len(Categories))
	for i := range Categories {
		descriptions = append(descriptions, Categories[i].Description)
	}
	bagSizes := []int{2, 3, 4}
	CM = closestmatch.New(descriptions, bagSizes)
	CMForCategories = make(map[string]*closestmatch.ClosestMatch)
	for i := range Categories {
		questions := make([]string, 0)
		for j := range Categories[i].QAs {
			questions = append(questions, j)
		}
		CMForCategories[descriptions[i]] = closestmatch.New(questions, bagSizes)
	}
	/*c1 := Category{
		Name:        "General",
		Description: "справи тебе привіт розумієш",
		QAs: map[string]string{
			"привіт":            "Привіт!",
			"як у тебе справи?": "Все чудово!",
			"ти мене розумієш?": "Я тебе розумію",
		},
	}
	categories = append(categories, c1)
	data, err := json.Marshal(&categories)
	if err != nil {
		loger.Log.Warningf("Cannot marshal matching data: %v", err)
	}
	if err = ioutil.WriteFile(config.Get().MatchingDataFile, data, 0666); err != nil {
		loger.Log.Warningf("Cannot write file: %v", err)
	}*/
	return nil
}
