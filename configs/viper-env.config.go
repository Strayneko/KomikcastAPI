package configs

import (
	"github.com/spf13/viper"
	"log"
)

var ViperEnv *viper.Viper

func InitViperEnvConfig() {
	ViperEnv = viper.New()
	ViperEnv.AddConfigPath("/")
	ViperEnv.SetConfigFile(".env")
	err := ViperEnv.ReadInConfig()

	if err != nil {
		log.Fatal("Failed to read env file" + err.Error())
	}
}
