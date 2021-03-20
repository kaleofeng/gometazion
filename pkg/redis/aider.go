package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Aider provides redis api wrapper.
type Aider struct {
	config Config
	db     redis.Conn
}

// NewAider new an instance.
func NewAider(config Config) *Aider {
	return &Aider{
		config: config,
		db:     nil,
	}
}

// Open open a connection.
func (aider *Aider) Open() error {
	rc := aider.config
	dsn := fmt.Sprintf("%s:%d", rc.Server, rc.Port)
	db, err := redis.Dial(rc.Network, dsn)
	if err != nil {
		return err
	}

	if _, err = db.Do("AUTH", rc.Password); nil != err {
		return err
	}

	if _, err = db.Do("SELECT", rc.Database); nil != err {
		return err
	}

	aider.db = db
	return nil
}

// Close close a connection.
func (aider *Aider) Close() error {
	return aider.db.Close()
}

// Do execute a command.
func (aider *Aider) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return aider.db.Do(commandName, args...)
}
