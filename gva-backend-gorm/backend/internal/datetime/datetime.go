package datetime

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nleeper/goment"
)

const FormatDateTime = "2006-01-02 15:04:05"
const FormatDate = "2006-01-02"

type (
	DateConfig struct {
		ctx         *context.Context
		location    *time.Location
		currentTime *goment.Goment
	}

	Option func(config *DateConfig)
)

type dateConfigKey struct{}

func (dc *DateConfig) Option(opts ...Option) {
	for _, opt := range opts {
		opt(dc)
	}
}

// wrap
func WrapContext(ctx context.Context, opts ...Option) context.Context {
	dateConfig := &DateConfig{}
	dateConfig.Option(opts...)
	return context.WithValue(ctx, dateConfigKey{}, dateConfig)
}

// extract
func GetConfigFromContext(ctx context.Context) (*DateConfig, error) {
	if now, ok := ctx.Value(dateConfigKey{}).(*DateConfig); ok {
		return now, nil
	}

	return nil, errors.New("DateConfig is not exist in context")
}

// Time creates a new option that sets the current time based on the given time value.
func Time(value time.Time) Option {
	return func(config *DateConfig) {
		t, err := goment.New(value.In(config.location))
		if err != nil {
			panic(fmt.Errorf("failed to create Goment from time: %w", err))

		}

		config.currentTime = t
	}
}

func GetNow(opts ...Option) (*goment.Goment, error) {
	config := &DateConfig{
		// default now is in hk timezone
		location: GetHongKongTimeLocation(),
	}

	config.Option(opts...)

	if config.ctx != nil {
		ctxConfig, err := GetConfigFromContext(*config.ctx)
		if err != nil {
			return nil, err
		}

		if ctxConfig.currentTime != nil {
			config.currentTime = ctxConfig.currentTime.Clone()
		}
	}

	if config.currentTime != nil {
		return config.currentTime, nil
	}

	now, err := goment.New(time.Now().In(config.location))
	if err != nil {
		return nil, fmt.Errorf("failed to get current time: %w", err)
	}

	return now, nil
}

func GetHongKongTimeLocation() *time.Location {
	return time.FixedZone("UTC+8", 8*3600)
}

func GetDateString(date string) (string, error) {
	config := &DateConfig{
		location: GetHongKongTimeLocation(),
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, config.location)

	formatedDate := parsedDate.Format(FormatDate)

	return formatedDate, nil
}

// get now from context if not now = now
func FromContext(ctx context.Context, opts ...Option) *goment.Goment {
	opts = append([]Option{WithContext(ctx)}, opts...)
	now, err := GetNow(opts...)
	if err != nil {
		panic(fmt.Errorf("MustNow failed: %w", err))
	}

	return now
}

func Must(opts ...Option) *goment.Goment {
	now, err := GetNow(opts...)
	if err != nil {
		panic(fmt.Errorf("MustNow failed: %w", err))
	}

	return now
}

func MustDateString(date string) string {
	new, err := GetDateString(date)
	if err != nil {
		panic(fmt.Sprintf("MustNew %v", err))
	}

	return new
}

/*
date: foramt "YYYY-MM-DD"
*/
func FromDateString(date string) (*goment.Goment, error) {
	dt, err := goment.New(date+" +0800", "YYYY-MM-DD ZZ")
	if err != nil {
		return nil, err
	}

	return dt, nil
}

/*
date: foramt "YYYY-MM-DD HH:mm:ss"

	"2024-07-05 00:00:00" -> "2024-07-05T08:00:00+08:00"

	"2024-07-05 23:59:59" -> "2024-07-06T07:59:59+08:00"
*/
func MustFromDateString(date string) *goment.Goment {
	dt, err := goment.New(date)
	dt.SetUTCOffset(8)

	if err != nil {
		panic(fmt.Sprintf("MustFromDateString %v", err))
	}

	return dt
}

// ChunkDateTime takes a date range and splits it into multiple ranges
// where each range is dayGap days apart
/*************  ✨ Codeium Command 🌟  *************/
func ChunkDateTime(ranges [2]*goment.Goment, dayGap int) [][2]*goment.Goment {
	start, end := ranges[0], ranges[1]
	chunks := make([][2]*goment.Goment, 0)

	if start.Clone().IsSame(end.Clone(), "day") {
		return [][2]*goment.Goment{ranges}
	}

	if end.Clone().IsBefore(start.Clone(), "second") {
		panic("ChunkDateTime end date is before start date")
	}

	maxGap := int(end.ToTime().Sub(start.ToTime()).Hours() / 24)
	if dayGap < maxGap {
		maxGap = dayGap
	}

	for i := start.Clone(); i.Clone().IsBefore(end.Clone(), "day"); {
		start := i.Clone().StartOf("day")
		next := i.Add(maxGap, "day").StartOf("day")

		endOfCurent := next.Clone().Add(-1, "second")
		if next.Clone().IsSame(end.Clone(), "day") {
			endOfCurent = end
		}

		chunks = append(chunks, [2]*goment.Goment{start, endOfCurent})
		i = next
	}

	return chunks
}
