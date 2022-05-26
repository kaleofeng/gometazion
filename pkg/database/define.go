package database

import "errors"

var (
	ErrorNoConfig            = errors.New("database: no conf file")
	ErrorUnsupportedDatabase = errors.New("database: unsupported database kind")
	ErrorOpenConnection      = errors.New("database: open connection failed")
	ErrorNoConnection        = errors.New("database: no connection available")
)
