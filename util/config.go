package util

import "github.com/spf13/viper"

type Config struct {
	ListenHost string `mapstructure:"LISTEN_HOST"`
	ListenPort string `mapstructure:"LISTEN_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env.development")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
