package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	NotionToken string
	NotionDatabaseID string
}

func Load() (*Config, error)  {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		NotionToken: viper.GetString("NOTION_TOKEN"),
		NotionDatabaseID: viper.GetString("NOTION_DATABASE_ID"),
	}

	return config, nil
}
