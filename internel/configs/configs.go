package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

type ForwarderConfig struct {
	ApiUrl       string
	FeeProvider  string
	EntryAddress string
	DiscordId    uint64
	DiscordToken string
	FeeTokens    map[string]int64
}

func LoadConfig() (*ForwarderConfig, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	config := &ForwarderConfig{}
	if err2 := json.Unmarshal(data, &config); err2 != nil {
		panic(err2)
	}
	{
		data, _ := json.Marshal(config)
		fmt.Println("config:", string(data))
	}
	return config, nil
}
