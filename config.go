package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	configFile string = "./config.toml"
)

type Config struct {
	Targets    []string `toml:"targets"`
	Extensions []string `toml:"extensions"`
}

func (c *Config) IsValid() bool {
	if len(c.Extensions) == 0 {
		return false
	}

	if len(c.Targets) == 0 {
		return false
	}

	return true
}

func (c *Config) PatternString() string {
	joined := strings.Join(c.Extensions, "|")
	pattern := `\.(` + joined + `)$`

	return pattern
}

func (c *Config) Extensions2Regexp() *regexp.Regexp {
	pattern := c.PatternString()
	return regexp.MustCompile(pattern)
}

func LoadConfig() (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	if !config.IsValid() {
		return nil, errors.New("Invalid toml file parameters")
	}

	return &config, nil
}
