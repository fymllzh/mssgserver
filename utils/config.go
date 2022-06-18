package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Viper *viper.Viper

func init() {

	dir, _:= os.Getwd()

	configFile := dir + "/conf/app.ini"


	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("config file %s not exists\n", configFile)
	}

	Viper = viper.New()
	Viper.SetConfigFile(configFile)
	Viper.SetConfigType("ini")
	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read config error: %v\n", err)
	}

	Viper.WatchConfig()
}

