package db

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormZapLogger struct {
	logger *zap.Logger
	// LogMode(LogLevel) Interface
	// Info(context.Context, string, ...interface{})
	// Warn(context.Context, string, ...interface{})
	// Error(context.Context, string, ...interface{})
	// Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

func (s GormZapLogger) LogMode(lvl logger.LogLevel) logger.Interface {
	return s
}
func (s GormZapLogger) Info(ctx context.Context, msg string, opt ...interface{}) {
	s.logger.Info("gorm", zap.String("msg", msg))
}
func (s GormZapLogger) Warn(ctx context.Context, msg string, opt ...interface{}) {
	s.logger.Warn("gorm", zap.String("msg", msg))
}
func (s GormZapLogger) Error(ctx context.Context, msg string, opt ...interface{}) {
	s.logger.Error("gorm", zap.String("msg", msg))
}
func (s GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
}

type InitOpt struct {
	Logger     *zap.Logger
	ConnectOpt struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string
		SSLMode  string
		TimeZone string
	}
}

func NewConnection(opt InitOpt) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s"+
			" port=%s sslmode=%s TimeZone=%s",
		opt.ConnectOpt.Host, opt.ConnectOpt.User, opt.ConnectOpt.Password,
		opt.ConnectOpt.DbName, opt.ConnectOpt.Port, opt.ConnectOpt.SSLMode, opt.ConnectOpt.TimeZone,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: GormZapLogger{
			logger: opt.Logger,
		},
	})

	return
}

type FoodItem struct {
	gorm.Model
	ID           *int
	Name         string   `json:",omitempty"`
	Description  string   `json:",omitempty"`
	Price        *float64 `json:",omitempty"`
	RestaurantID *int     `json:",omitempty"`
	// Restaurant   Restaurant
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`
}

func Init(opt InitOpt) (db *gorm.DB, err error) {
	db, err = NewConnection(opt)
	if err != nil {
		return
	}

	err = db.AutoMigrate(&FoodItem{})
	if err != nil {
		return
	}
	return
}
