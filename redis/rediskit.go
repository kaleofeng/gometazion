package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// RedisKit provides redis api wrapper.
type RedisKit struct {
	redisConfig RedisConfig
	db          redis.Conn
}

// New new an object.
func (redisKit *RedisKit) New(redisConfig RedisConfig) *RedisKit {
	fmt.Printf("Redis Kit - New: config(%v)\n", redisKit.redisConfig)

	redisKit.redisConfig = redisConfig
	return redisKit
}

// Open open a connection.
func (redisKit *RedisKit) Open() bool {
	fmt.Printf("Redis Kit - Open: config(%v)\n", redisKit.redisConfig)

	rc := redisKit.redisConfig
	dsn := fmt.Sprintf("%s:%d", rc.Server, rc.Port)
	db, err := redis.Dial(rc.Network, dsn)
	if err != nil {
		fmt.Printf("Redis Kit - Open: redis dial failed, err(%v)\n", err)
		return false
	}

	if _, err = db.Do("AUTH", rc.Password); nil != err {
		fmt.Printf("Redis Kit - Open: db auth failed, err(%v)\n", err)
		return false
	}

	if _, err = db.Do("SELECT", rc.Database); nil != err {
		fmt.Printf("Redis Kit - Open: db select failed, err(%v)\n", err)
		return false
	}

	redisKit.db = db
	fmt.Printf("Redis Kit - Open: success, config(%v)\n", redisKit.redisConfig)
	return true
}

// Close close a connection.
func (redisKit *RedisKit) Close() {
	redisKit.db.Close()
	fmt.Printf("Redis Kit - Close: config(%v)\n", redisKit.redisConfig)
}

// Do execute a command.
func (redisKit *RedisKit) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return redisKit.db.Do(commandName, args...)
}
