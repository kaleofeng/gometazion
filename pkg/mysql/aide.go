package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

// Aide provides mysql api wrapper.
type Aide struct {
	config Config
	db     *sql.DB
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
	mc := aide.config
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

	aide.db = db
	return nil
}

// Close close a connection.
func (aide *Aide) Close() error {
	return aide.db.Close()
}

// Query query rows.
func (aide *Aide) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return aide.db.Query(query, args...)
}
