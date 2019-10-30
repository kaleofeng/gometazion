package mz

import (
	"database/sql"
	"fmt"
	"time"
)

// MysqlKit provides mysql api wrapper.
type MysqlKit struct {
	mysqlConfig MysqlConfig
	db          *sql.DB
}

// New an object.
func (mysqlKit *MysqlKit) New(mysqlConfig MysqlConfig) *MysqlKit {
	mysqlKit.mysqlConfig = mysqlConfig
	return mysqlKit
}

// Open a connection.
func (mysqlKit *MysqlKit) Open() bool {
	fmt.Printf("Mysql Kit open: %v\n", mysqlKit.mysqlConfig)

	mc := mysqlKit.mysqlConfig
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mc.Username, mc.Password, mc.Network, mc.Server, mc.Port, mc.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Mysql Kit open failed, err: %v\n", err)
		return false
	}

	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	if err = db.Ping(); nil != err {
		fmt.Printf("Mysql Kit ping failed, err: %v\n", err)
		return false
	}

	mysqlKit.db = db
	fmt.Printf("Mysql Kit open success: %v\n", mysqlKit.mysqlConfig)
	return true
}

// Query rows.
func (mysqlKit *MysqlKit) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return mysqlKit.db.Query(query, args...)
}
