package matching

import (
	"encoding/json"
	"io/ioutil"

	"github.com/KitlerUA/NNCompanionBot/config"
	"github.com/KitlerUA/NNCompanionBot/loger"
	"github.com/schollz/closestmatch"
)

var CMGeneral *closestmatch.ClosestMatch
var QAGeneral map[string]string

func PrepareDataForMatching() error {
	data, err := ioutil.ReadFile(config.Get().MatchingDataFile)
	if err != nil {
		loger.Log.Warningf("Cannot find file with matching data: %v", err)
		return err
	}
	QAGeneral = make(map[string]string)
	if err = json.Unmarshal(data, &QAGeneral); err != nil {
		loger.Log.Warningf("Corrupted data in matching file: %v", err)
		return err
	}
	generalListOfQ := make([]string, 0, len(QAGeneral))
	bagSizes := []int{3, 4, 5}
	for i := range QAGeneral {
		generalListOfQ = append(generalListOfQ, i)
	}
	CMGeneral = closestmatch.New(generalListOfQ, bagSizes)
	return nil
}
