package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

// Aider provides mysql api wrapper.
type Aider struct {
	config Config
	db     *sql.DB
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
	mc := aider.config
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mc.Username, mc.Password, mc.Network, mc.Server, mc.Port, mc.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	if err = db.Ping(); nil != err {
		return err
	}

	aider.db = db
	return nil
}

// Close close a connection.
func (aider *Aider) Close() error {
	return aider.db.Close()
}

// Query query rows.
func (aider *Aider) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return aider.db.Query(query, args...)
}
