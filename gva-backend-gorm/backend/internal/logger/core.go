package logger

import (
	"backend/env"
	"backend/internal/logger/report"
	"context"

	"runtime/debug"

	"github.com/tidwall/pretty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Core struct {
	zapcore.Core
	env *env.Config
}

func NewCore(zapCore zapcore.Core, env *env.Config) *Core {
	return &Core{zapCore, env}
}

func (c *Core) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checked.AddCore(entry, c)
	}
	return checked
}

func WithLazy(log *zap.Logger, fields ...zap.Field) *zap.Logger {
	if len(fields) == 0 {
		return log
	}
	return log.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewLazyWith(NewCore(core, nil), fields)
	}))
}

var encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	// Keys can be anything except the empty string.
	NameKey:        "logName",
	CallerKey:      "logCaller",
	FunctionKey:    zapcore.OmitKey,
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
})

func PretifyZapFields(entry zapcore.Entry, fields []zapcore.Field) []byte {
	buf, _ := encoder.EncodeEntry(entry, fields)
	return pretty.Pretty(buf.Bytes())
}

func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.LoggerName == "gorm" && entry.Level == zap.ErrorLevel {
		// ignore beside slow query
		return c.Core.Write(entry, fields)
	}

	switch entry.Level {
	case zap.ErrorLevel, zap.WarnLevel, zap.FatalLevel, zap.PanicLevel, zap.InfoLevel:
		icon := "🔥"
		tag := "error"

		switch entry.Level {
		case zap.WarnLevel:
			icon = "⚠️"
			tag = "warn"
		case zap.FatalLevel:
			tag = "fetal"
		}

		go report.Send(context.Background(),
			report.WithTitle(entry.Message),
			report.WithIcon(icon),
			report.AddTag(tag),
			report.AddMessage("", string(PretifyZapFields(entry, fields))),
			report.AddMessage("Stack", string(debug.Stack())),
			report.WithUrl(c.env.Google.ChatWebhookURL),
			report.WithMode(c.env.App.Env),
		)
	}

	return c.Core.Write(entry, fields)
}
