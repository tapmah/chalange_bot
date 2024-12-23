package config

import (
	"os"
)

type Config struct {
	Token       string
	AdminChatId string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
}

var globalConfig *Config

func GetConfig() *Config {
	if globalConfig == nil {
		LoadConfig()
	}
	return globalConfig
}

func LoadConfig() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	globalConfig = &Config{
		Token:      token,
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBName:     os.Getenv("DBName"),
	}
}
