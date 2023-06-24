package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type BotConfig struct {
	AuthFile string   `yaml:"bot_auth_file"`
	LogLevel string   `yaml:"bot_log_level"`
	Prefixes []string `yaml:"bot_prefixes"`
}

func InitConfig(config_path string) error {
	bytes, err := os.ReadFile(config_path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(bytes, BotConfigVar); err != nil {
		return err
	}

	if len(BotConfigVar.Prefixes) < 1 {
		BotConfigVar.Prefixes = []string{"$$", "!"}
	} else if len(BotConfigVar.LogLevel) < 1 {
		BotConfigVar.LogLevel = "DEBUG"
	}

	return err
}
