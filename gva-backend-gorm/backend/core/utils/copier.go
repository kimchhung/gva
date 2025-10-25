package utils

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func MustCopy[T any](toValue T, from any) T {
	if err := copier.Copy(toValue, from); err != nil {
		panic(fmt.Errorf("MustCopy err %v", err))
	}

	return toValue
}
