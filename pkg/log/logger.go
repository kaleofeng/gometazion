package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const (
	PanicLevel logrus.Level = logrus.PanicLevel
	FatalLevel              = logrus.FatalLevel
	ErrorLevel              = logrus.ErrorLevel
	WarnLevel               = logrus.WarnLevel
	InfoLevel               = logrus.InfoLevel
	DebugLevel              = logrus.DebugLevel
	TraceLevel              = logrus.TraceLevel
)

type Config struct {
	Dir      string
	Filename string
	Level    logrus.Level
}

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(config *Config) (*Logger, error) {
	err := os.MkdirAll(config.Dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(config.Dir, config.Filename)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.Out = io.MultiWriter(os.Stdout, file)
	logger.SetLevel(config.Level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		logger,
	}, nil
}

func (logger *Logger) Panicf(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Panicf(format, args...)
}

func (logger *Logger) Fatalf(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Fatalf(format, args...)
}

func (logger *Logger) Errorf(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Errorf(format, args...)
}

func (logger *Logger) Warnf(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Warnf(format, args...)
}

func (logger *Logger) Infof(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Infof(format, args...)
}

func (logger *Logger) Debugf(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Debugf(format, args...)
}

func (logger *Logger) Tracef(tag string, format string, args ...interface{}) {
	logger.logger.WithField("tag", tag).Tracef(format, args...)
}
