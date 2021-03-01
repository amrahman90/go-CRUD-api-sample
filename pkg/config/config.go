package config

import (
	"github.com/amrahman90/go-CRUD-api-sample/pkg/db"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Env  string
		Host string
		Port string
	}

	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string
		SSLMode  string
		TimeZone string
	}

	Logger struct {
		OutputPath        string
		Level             string
		DisableStackTrace bool
	}
	logger *zap.Logger
	dbConn *gorm.DB
}

func (c Config) GetLogger() *zap.Logger {
	if c.logger != nil {
		return c.logger
	}
	log := logger.InitLogger(logger.Config{
		LogLevel:          c.Logger.Level,
		LogOutputPaths:    c.Logger.OutputPath,
		DisableStackTrace: c.Logger.DisableStackTrace,
	})
	c.logger = log
	return log
}

func (c Config) GetDB() (conn *gorm.DB, err error) {
	if c.dbConn != nil {
		return c.dbConn, nil
	}
	conn, err = db.Init(db.InitOpt{
		Logger:     c.GetLogger(),
		ConnectOpt: c.Db,
	})
	if err != nil {
		return
	}
	c.dbConn = conn
	return
}
