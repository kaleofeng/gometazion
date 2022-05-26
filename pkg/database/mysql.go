package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openMysql(conf *Config, gormConfig *gorm.Config) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(format, conf.Username, conf.Password, conf.Host, conf.Port, conf.DB, conf.Charset)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
