package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Aide provides redis api wrapper.
type Aide struct {
	config Config
	db     redis.Conn
}

// NewAide new an instance.
func NewAide(config Config) *Aide {
	return &Aide{
		config: config,
		db:     nil,
	}
}

// Open open a connection.
func (aide *Aide) Open() error {
	rc := aide.config
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

	aide.db = db
	return nil
}

// Close close a connection.
func (aide *Aide) Close() error {
	return aide.db.Close()
}

// Do execute a command.
func (aide *Aide) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return aide.db.Do(commandName, args...)
}
