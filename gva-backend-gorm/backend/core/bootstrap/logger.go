package bootstrap

import (
	"backend/core/env"
	"backend/internal/report"
	"context"
	"fmt"
	"runtime/debug"

	prettyconsole "github.com/thessem/zap-prettyconsole"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
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
)

func NewZapLogger(core *Logger) *zap.Logger {
	if err := core.Initailized(); err != nil {
		panic(err)
	}

	return core.logger
}

type Logger struct {
	zapcore.Core
	env    *env.Config
	logger *zap.Logger
}

func NewLogger(env *env.Config) *Logger {
	l := &Logger{
		env: env,
	}
	return l
}

func (c *Logger) Logger() *zap.Logger {
	return c.logger
}

func (c *Logger) Initailized() (err error) {
	if c.logger != nil {
		return nil
	}

	var config zap.Config
	if c.env.Logger.Prettier {
		config = prettyconsole.NewConfig()
	} else {
		if c.env.IsProd() {
			config = zap.NewProductionConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}
	}

	config.Level = zap.NewAtomicLevelAt(zapcore.Level(c.env.Logger.Level))
	c.logger, err = config.Build(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.WrapCore(c.CustomCore()),
	)

	if err != nil {
		return fmt.Errorf("failed to build logger config %v", err)
	}

	zap.ReplaceGlobals(c.logger)
	return nil
}

func (c *Logger) CustomCore() func(zapcore.Core) zapcore.Core {
	return func(core zapcore.Core) zapcore.Core {
		c.Core = core
		return c
	}
}

func (c *Logger) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Core.Enabled(entry.Level) {
		return checked.AddCore(entry, c.Core)
	}
	return checked
}

func (c *Logger) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.LoggerName == "gorm" && entry.Level == zap.ErrorLevel {
		// ignore beside slow query
		return c.Core.Write(entry, fields)
	}

	switch entry.Level {
	case zap.ErrorLevel, zap.WarnLevel, zap.FatalLevel, zap.PanicLevel:
		c.sendReportQueue(entry, fields)
	}

	return c.Core.Write(entry, fields)
}

func (c *Logger) sendReportQueue(entry zapcore.Entry, fields []zapcore.Field) {
	if !c.env.Google.Enable {
		return
	}

	icon := "🔥"
	tag := "error"

	switch entry.Level {
	case zap.WarnLevel:
		icon = "⚠️"
		tag = "warn"
	case zap.FatalLevel:
		tag = "fetal"
	case zap.PanicLevel:
		tag = "panic"
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

func PretifyZapFields(entry zapcore.Entry, fields []zapcore.Field) []byte {
	buf, _ := encoder.EncodeEntry(entry, fields)
	return pretty.Pretty(buf.Bytes())
}
