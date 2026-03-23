package config

import (
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	slog.InfoContext(ctx, msg, "data", data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	slog.WarnContext(ctx, msg, "data", data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	slog.ErrorContext(ctx, msg, "data", data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		slog.ErrorContext(ctx, "GORM Query Error",
			"err", err,
			"elapsed", elapsed.String(),
			"rows", rows,
			"sql", sql,
		)
		return
	}

	if l.LogLevel == logger.Info {
		slog.InfoContext(ctx, "GORM Query",
			"elapsed", elapsed.String(),
			"rows", rows,
			"sql", sql,
		)
	}
}
