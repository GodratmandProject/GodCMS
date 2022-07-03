package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var ErrUnknownFileExtension = errors.New("config file not supported")

func Parse(cfg *Config, path string) error {
	switch filepath.Ext(path) {
	case ".yml", ".yaml":
		return parseYaml(cfg, path)

	case ".json":
		return parseJson(cfg, path)
	default:
		return ErrUnknownFileExtension
	}
}

func parseJson(cfg *Config, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	if err = json.NewDecoder(f).Decode(cfg); err != nil {
		return err
	}
	return nil
}

func parseYaml(cfg *Config, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	if err = yaml.NewDecoder(f).Decode(cfg); err != nil {
		return err
	}
	return nil
}