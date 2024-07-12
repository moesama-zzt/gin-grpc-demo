package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
