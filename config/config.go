package config

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed words.yaml
var wordsConfig []byte

type Config struct {
	Words []Word `yaml:"words"`
}

type Word struct {
	Word string `yaml:"word"`
	Hint string `yaml:"hint"`
}

func Init(path *string) (*Config, error) {
	var err error
	if path != nil && *path != "" {
		wordsConfig, err = os.ReadFile(*path)
		if err != nil {
			return nil, err
		}
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(wordsConfig, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
