package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"log/slog"
	"strings"
)

type LogrusHandler struct {
	logger *logrus.Logger
}

func NewLogrusHandler(logger *logrus.Logger) *LogrusHandler {
	return &LogrusHandler{logger: logger}
}

func ConvertLogLevel(level string) logrus.Level {
	var l logrus.Level
	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warm":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}
	return l
}

func (h *LogrusHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// support all logging levels
	return true
}

func (h *LogrusHandler) Handle(ctx context.Context, rec slog.Record) error {
	fields := make(map[string]interface{}, rec.NumAttrs())

	rec.Attrs(func(attr slog.Attr) bool {
		fields[attr.Key] = attr.Value.Any()
		return true
	})

	entry := h.logger.WithFields(fields)

	switch rec.Level {
	case slog.LevelDebug:
		entry.Debug(rec.Message)
	case slog.LevelInfo:
		entry.Info(rec.Message)
	case slog.LevelWarn:
		entry.Warn(rec.Message)
	case slog.LevelError:
		entry.Error(rec.Message)
	}

	return nil
}

func (h *LogrusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// not implemented for brevity
	return h
}

func (h *LogrusHandler) WithGroup(name string) slog.Handler {
	// not implemented for brevity
	return h
}
