package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	BindPort string
}

type DatabaseConfig struct {
	Scheme       string
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	SslMode      string
}

type MetricsConfig struct {
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Metrics  MetricsConfig
}

func InitializeConfig(configFilePath string, configFileName string) (*Config, error) {

	var config *Config

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		// TODO log to stderr as main logger has not been initialized yet.
		log.Fatalf("error in reading config - %v", err)
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
