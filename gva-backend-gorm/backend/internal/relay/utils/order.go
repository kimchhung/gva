package utils

import (
	"strconv"
	"strings"
)

func AppendOrder(acc, field, direction string) string {
	if acc == "" {
		return field + " " + direction
	}

	return acc + ", " + field + " " + direction
}

func ReverseDirection(direction string) string {
	if direction == "ASC" || direction == "asc" {
		return "DESC"
	}

	return "ASC"
}

func SplitOrderKey(key string) (int64, string) {
	if !strings.Contains(key, "-") {
		return -1, key
	}

	parts := strings.SplitN(key, "-", 2)
	order, field := parts[0], parts[1]
	i, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		panic("order key without order number")
	}

	return i, field
}
