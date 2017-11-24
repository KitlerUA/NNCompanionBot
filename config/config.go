package config

import (
	"sync"
	"io/ioutil"
	"github.com/KitlerUA/NNCompanionBot/loger"
	"encoding/json"
)

const defaultConfigFilename = "config.json"

type Config struct {
	Token string `json:"token"`
}

var config *Config
var once sync.Once

func GetConfig() Config{
	once.Do(readConfigFile)
	return *config
}

func readConfigFile() {
	data, err := ioutil.ReadFile(defaultConfigFilename)
	if err != nil {
		loger.Log.Panicf("Cannot read config file: %v", err)
	}
	config = &Config{}
	if err = json.Unmarshal(data, &config); err != nil {
		loger.Log.Panicf("Corrupted data in config file: &v", err)
	}
}