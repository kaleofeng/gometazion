package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Service struct {
	conf Config
	gdb  *gorm.DB
}

func (s *Service) Start(confPath string) error {
	if err := s.conf.Load(confPath); err != nil {
		return err
	}

	if err := s.openConnection(); err != nil {
		return ErrorOpenConnection
	}

	return nil
}

func (s *Service) DB() (*gorm.DB, error) {
	if s.gdb == nil {
		return nil, ErrorNoConnection
	}
	return s.gdb, nil
}

func (s *Service) openConnection() error {
	var gdb *gorm.DB
	var err error

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}

	switch s.conf.Kind {
	case "mysql":
		gdb, err = openMysql(&s.conf, &gormConfig)
	default:
		err = ErrorUnsupportedDatabase
	}

	if err != nil {
		return err
	}

	s.gdb = gdb
	return nil
}
