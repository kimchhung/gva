package utils

import (
	"backend/internal/gormq"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"backend/core/utils/json"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nleeper/goment"
	"golang.org/x/exp/rand"
)

func IsEnabled(key bool) func(c echo.Context) bool {
	if key {
		return nil
	}

	return func(c echo.Context) bool { return true }
}

func SetIfEmpty[T comparable](dest *T, value T) {
	if dest == nil {
		dest = new(T)
	}

	if IsEmpty(dest) {
		dest = &value
	}
}

func IsEmpty[T comparable](v T) bool {
	var zero T
	return v == zero
}

func PanicIfErr(prefix string, err error) {
	if err == nil {
		return
	}
	panic(fmt.Errorf("%s %v", prefix, err))
}

func Async[A any](f func() A) <-chan A {
	ch := make(chan A, 1)
	go func() {
		ch <- f()
	}()
	return ch
}

func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}
	return *x
}

func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

func FromPtrOr[T any](pointer *T, fallback T) T {
	if pointer == nil {
		return fallback
	}
	return *pointer
}

func ToPtr[T any](x T) *T {
	return &x
}

func Empty[T any]() T {
	var zero T
	return zero
}

func StructToMap(p any, columnMap gormq.ColumnMap) (dbCols map[string]any, resp map[string]any) {
	dbCols = make(map[string]any)
	resp = make(map[string]any)

	// Use reflection to populate the payload with non-nil values from p
	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldName := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]

		// Check if the field is a pointer and not nil
		if !fieldValue.IsNil() {
			// Convert field name to lowercase and add to the payload
			if column, ok := columnMap[fieldName]; ok {
				dbCols[column] = fieldValue.Interface()
				resp[fieldName] = fieldValue.Interface()
				switch fieldValue.Kind() {
				case reflect.Slice:
					dbCols[column] = json.MustJSON(fieldValue.Interface())
				default:
					dbCols[column] = fieldValue.Interface()
				}
			}
		}
	}

	return dbCols, resp
}

func GetMapKeys(m map[string]any) (keys []string) {
	for key := range m {
		return append(keys, key)
	}
	return
}

func GoForEach[T any](collection []T, callback func(item T, index int)) {
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			callback(_item, _i)
			wg.Done()
		}(item, i)
	}

	wg.Wait()
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func SleepRandom(max time.Duration) {
	delay := Random(0, int(max))
	time.Sleep(time.Duration(delay))
}

func HashX(rawDate string, result []int) string {
	resultStr := strings.Join(Map(result, func(item, _ int) string {
		return IntToStr(item)
	}), ",")

	return Hash(rawDate + ":" + resultStr)
}

func IntToStr(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func StrToInt(value string) int {
	vInt, _ := strconv.Atoi(value)
	return vInt
}

func Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Map[T any, R any](collection []T, callback func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = callback(item, i)
	}

	return result
}

func Filter[V any](collection []V, callback func(item V, index int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if callback(item, i) {
			result = append(result, item)
		}
	}

	return result
}

func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result := []R{}

	for i, item := range collection {
		if r, ok := callback(item, i); ok {
			result = append(result, r)
		}
	}

	return result
}

func FromTimeStr(str string, now *goment.Goment) (todayTime *goment.Goment) {
	hour, minute, second := 0, 0, 0
	for i, text := range strings.Split(str, ":") {
		switch i {
		case 0:
			hour = StrToInt(text)
		case 1:
			minute = StrToInt(text)
		case 2:
			second = StrToInt(text)
		}
	}

	todayTime = now.Clone().SetHour(hour).SetMinute(minute).SetSecond(second)
	return
}

type AttemptConfig[T any] struct {
	Attempts    int
	Delay       time.Duration
	DelayFn     func(i int) time.Duration
	TryOnError  func(result T, err error) bool
	IsSaveStats bool
}

func AttemptX[T any](ctx context.Context, config *AttemptConfig[T], f func(ctx context.Context, i int) (T, error)) (result T, err error) {
	forceStopRetry := false
	delay := config.Delay

	if config.Attempts < 1 {
		panic("Attempts must be > 0")
	}

	for i := 0; i < config.Attempts; i++ {
		select {
		case <-ctx.Done():
			return result, fmt.Errorf("context cancel, stop attempting %v", ctx.Err())
		default:
			if forceStopRetry {
				return result, fmt.Errorf("after %d Attempts, last err: %s", config.Attempts, err)
			}

			result, err = f(ctx, i)
			if err == nil {
				return result, nil
			}

			if config.DelayFn != nil {
				delay = config.DelayFn(i)
			}

			if config.TryOnError != nil {
				forceStopRetry = !config.TryOnError(result, err)
			}

			time.Sleep(delay)
		}
	}

	return result, fmt.Errorf("after %d Attempts, last err: %s", config.Attempts, err)
}

func Attempts[T any](ctx context.Context, attempts int, sleep time.Duration, f func(ctx context.Context, i int) (T, error)) (result T, err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			time.Sleep(sleep)
		}

		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
			result, err = f(ctx, i)
			if err == nil {
				return result, nil
			}
		}
	}
	return result, fmt.Errorf("after %d Attempts, last err: %s", attempts, err)
}
