package data

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	gormLogger "gorm.io/gorm/logger"
)

type GormLoggerHelper struct {
	msgKey  string
	logger  log.Logger
	logMode log.Level
}

func (g *GormLoggerHelper) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	switch level {
	case gormLogger.Silent:
		g.logMode = log.LevelFatal
	case gormLogger.Error:
		g.logMode = log.LevelError
	case gormLogger.Info:
		g.logMode = log.LevelInfo
	case gormLogger.Warn:
		g.logMode = log.LevelWarn
	}
	return g
}

func (g *GormLoggerHelper) Info(ctx context.Context, msg string, vals ...interface{}) {
	if g.logMode <= log.LevelInfo {
		g.logger.Log(log.LevelInfo, g.msgKey, fmt.Sprintf(msg, vals...), "location", FileWithLineNum())
	}
}

func (g *GormLoggerHelper) Warn(ctx context.Context, msg string, vals ...interface{}) {
	if g.logMode <= log.LevelWarn {
		g.logger.Log(log.LevelWarn, g.msgKey, fmt.Sprintf(msg, vals...), "location", FileWithLineNum())
	}
}

func (g *GormLoggerHelper) Error(ctx context.Context, msg string, vals ...interface{}) {
	if g.logMode <= log.LevelError {
		g.logger.Log(log.LevelError, g.msgKey, fmt.Sprintf(msg, vals...), "location", FileWithLineNum())
	}
}

func (g *GormLoggerHelper) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	switch g.logMode {
	case log.LevelDebug, log.LevelInfo, log.LevelWarn:
		if err != nil {
			g.logger.Log(log.LevelError, g.msgKey, err)
		}
		g.logger.Log(log.LevelInfo, g.msgKey, fmt.Sprintf("%s\nrowAffected: %d\n", sql, rowsAffected))
	case log.LevelError, log.LevelFatal:
		if err != nil {
			g.logger.Log(log.LevelError, g.msgKey, err)
			g.logger.Log(log.LevelError, g.msgKey, fmt.Sprintf("%s\nrowAffected: %d\n", sql, rowsAffected))
		}
	}
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}

func NewGormLoggerHelper(log log.Logger, level gormLogger.LogLevel) gormLogger.Interface {
	logger := GormLoggerHelper{logger: log, msgKey: "msg"}
	return logger.LogMode(level)
}
