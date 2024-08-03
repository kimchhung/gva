package utils

import (
	"errors"
	"time"
)

type RetryFunc func() error

func Retry(retryFunc RetryFunc, maxIterations int, delay time.Duration) error {
	var err error
	for i := 0; i < maxIterations; i += 1 {
		func() {
			defer func() {
				if r := recover(); r != nil {
					switch recoverType := r.(type) {
					case string:
						err = errors.New(recoverType)
					case error:
						err = recoverType
					default:
						err = errors.New("unexpected type")
					}
				}
			}()
			err = retryFunc()
		}()
		if err == nil {
			return nil
		}
		if i < maxIterations {
			time.Sleep(delay)
		}
	}
	return err
}
