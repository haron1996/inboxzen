package viper

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBString     string `mapstructure:"dbString"`
	PORT         string `mapstructure:"port"`
	SecretKeyHex string `mapstructure:"secretKeyHex"`
	PublicKeyHex string `mapstructure:"publicKeyHex"`
	HoldLabel    string `mapstructure:"holdLabel"`
	BlockLabel   string `mapstructure:"blockLabel"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
