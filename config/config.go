package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	BindPort string
}

func InitializeConfig(configFilePath string, configFileName string) (*Config, error) {

	var config *Config

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		// TODO log to stderr as main logger has not been initialized yet.
		fmt.Errorf("error in reading config - %v", err)
		return nil, err
	}
	return unmarshallConfig(config)
}

func unmarshallConfig(config *Config) (*Config, error) {
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Errorf("error in unmarshalling config file")
		return nil, err
	}
	return config, nil

}
