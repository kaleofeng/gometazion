package mysql

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents mysql configs.
type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// Load load mysql configs.
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
