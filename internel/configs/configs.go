package configs

import (
	"github.com/spf13/viper"
)

var Viper *viper.Viper

func LoadConfig() (*viper.Viper, error) {
	Viper = viper.New()
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return Viper, nil
}
