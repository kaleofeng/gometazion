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
	fmt.Printf("Mysql Kit - New: config(%v)\n", mysqlKit.mysqlConfig)

	mysqlKit.mysqlConfig = mysqlConfig
	return mysqlKit
}

// Open a connection.
func (mysqlKit *MysqlKit) Open() bool {
	fmt.Printf("Mysql Kit - Open: config(%v)\n", mysqlKit.mysqlConfig)

	mc := mysqlKit.mysqlConfig
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mc.Username, mc.Password, mc.Network, mc.Server, mc.Port, mc.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Mysql Kit - Open: sql open failed, err(%v)\n", err)
		return false
	}

	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	if err = db.Ping(); nil != err {
		fmt.Printf("Mysql Kit - Open: db ping failed, err(%v)\n", err)
		return false
	}

	mysqlKit.db = db
	fmt.Printf("Mysql Kit - Open: success, config(%v)\n", mysqlKit.mysqlConfig)
	return true
}

// Close a connection.
func (mysqlKit *MysqlKit) Close() {
	mysqlKit.db.Close()
	fmt.Printf("Mysql Kit - Close: config(%v)\n", mysqlKit.mysqlConfig)
}

// Query rows.
func (mysqlKit *MysqlKit) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return mysqlKit.db.Query(query, args...)
}
