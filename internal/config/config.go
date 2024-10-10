package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JiraToken    string `mapstructure:"JIRA_TOKEN"`
	JiraUsername string `mapstructure:"JIRA_USERNAME"`
	JiraHost     string `mapstructure:"JIRA_HOST"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
