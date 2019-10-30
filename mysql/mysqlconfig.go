package mz

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// MysqlConfig represents a db config
type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// Load mysql configs.
func (mysqlConfig *MysqlConfig) Load() bool {
	yamlFile, err := ioutil.ReadFile("mysql.yaml")
	if err != nil {
		fmt.Printf("Mysql Config load failed, err: %v\n", err)
		return false
	}

	err = yaml.Unmarshal(yamlFile, mysqlConfig)
	if err != nil {
		fmt.Printf("Mysql Config unmarshal failed, err: %v\n", err)
		return false
	}

	fmt.Printf("Mysql Config load success: %v\n", mysqlConfig)
	return true
}
