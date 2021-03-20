package redis

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents redis config
type Config struct {
	Password string `yaml:"password"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database int    `yaml:"database"`
}

// Load load redis configs.
func (config *Config) Load(filePath string) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}

	return nil
}
