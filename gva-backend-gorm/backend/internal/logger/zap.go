package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"backend/env"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ *zap.Logger

// G stand for Global, clone from with using logger.G().With
func G() *zap.Logger {
	return zap.L()
}

func nameFromStruct(from any) string {
	name := ""
	switch reflect.TypeOf(from).Kind() {
	case reflect.Ptr:
		name = reflect.TypeOf(from).Elem().Name()
	case reflect.String:
		name = from.(string)
	case reflect.Struct:
		name = reflect.TypeOf(from).Name()
	default:
		panic("invalid from params")
	}

	return name
}

// clone from global and add where it came from
func Clone(from interface{}) *zap.Logger {
	return G().With().Named(nameFromStruct(from))
}

func WithCore(env *env.Config) zap.Option {
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return NewCore(c, env)
	})
}

func Log(v ...any) {
	if b, ok := v[0].([]byte); ok {
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, b, "", "   "); err != nil {
			panic(err)
		}

		fmt.Println("----------------- json --------------")
		fmt.Println(dst.String())
	} else {
		s, _ := json.MarshalIndent(v, "", "   ")
		fmt.Println("----------------- json --------------")
		fmt.Println(string(s))
	}
}

// recover panic crash with report message
func Recover(fields ...zapcore.Field) {
	if r := recover(); r != nil {
		fields = append(fields, zap.Any("err", r))
		G().Error("panic recover",
			fields...,
		)
	}
}
