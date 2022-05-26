package database

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/kaleofeng/gometazion/pkg/kit/file"
)

type Config struct {
	Kind     string `yaml:"Kind"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DB       string `yaml:"DB"`
	Charset  string `yaml:"Charset"`
}

func (c *Config) Load(confPath string) error {
	if exists, _ := file.Exists(confPath); !exists {
		return ErrorNoConfig
	}

	bytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	var conf Config
	if err = yaml.Unmarshal(bytes, &conf); err != nil {
		return err
	}

	*c = conf
	return nil
}

func (c *Config) Save(confPath string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	confDir := filepath.Dir(confPath)
	if err := os.MkdirAll(confDir, 0600); err != nil {
		return err
	}

	if err = ioutil.WriteFile(confPath, bytes, 0600); err != nil {
		return err
	}

	return nil
}
