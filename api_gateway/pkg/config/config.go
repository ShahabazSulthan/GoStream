package config

import "github.com/spf13/viper"

type Config struct {
	VideoService string `mapstructure:"VIDEO_SRV"`
	Port         string `mapstructure:"PORT"`
}

var envs = []string{"VIDEO_SRV", "PORT"}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}