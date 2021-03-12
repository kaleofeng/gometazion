package redis

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// RedisConfig represents a db config
type RedisConfig struct {
	Password string `yaml:"password"`
	Network  string `yaml:"network"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// Load load redis configs.
func (redisConfig *RedisConfig) Load() bool {
	fmt.Printf("Redis Config - Load: config(%v)\n", redisConfig)

	yamlFile, err := ioutil.ReadFile("redis.yaml")
	if err != nil {
		fmt.Printf("Redis Config - Load: read file failed, err(%v)\n", err)
		return false
	}

	err = yaml.Unmarshal(yamlFile, redisConfig)
	if err != nil {
		fmt.Printf("Redis Config - Load: unmarshal failed, err(%v)\n", err)
		return false
	}

	fmt.Printf("Redis Config - Load: success, config(%v)\n", redisConfig)
	return true
}
