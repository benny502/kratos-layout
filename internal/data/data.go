package data

import (
	"os"
	"time"

	golog "log"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserData)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db := Mysql(c.Mysql, logger)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func (d *Data) GetDB() *gorm.DB {
	return d.db
}

// Mysql
func Mysql(cfg *conf.Data_Mysql, logger log.Logger) *gorm.DB {
	logMode := gormLogger.Error
	loggerInterface := NewGormLoggerHelper(logger, logMode)
	if cfg.LogMode == "debug" {
		logMode = gormLogger.Info
		loggerInterface = gormLogger.New(golog.New(os.Stdout, "\r\n", golog.LstdFlags), gormLogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logMode,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}) // 将默认打印sql级别调整为info
	}
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		PrepareStmt: cfg.PrepareStmt,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,
			SingularTable: cfg.SingularTable,
		},
		Logger: loggerInterface,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(int(cfg.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(cfg.MaxOpenConns))
	sqlDB.SetConnMaxLifetime(time.Duration(int64(cfg.ConnMaxLifetime)) * time.Second)
	return db
}
