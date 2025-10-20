package cache

import (
	"fmt"
	"strings"
	"time"
)

const (
	AGE_30_MIN    = time.Hour / 2
	AGE_ONE_HOUR  = time.Hour
	AGE_FOUR_HOUR = time.Hour * 4
	AGE_HALF_DAY  = time.Hour * 12
	AGE_ONE_DAY   = time.Hour * 24
	AGE_ONE_WEEK  = time.Hour * 24 * 7
	AGE_ONE_MONTH = time.Hour * 24 * 7 * 30
)

// RedisKey is a type that represents a key in Redis.
type RedisKey string

// String returns the string representation of the RedisKey.
func (k RedisKey) String() string {
	return string(k)
}

func (k RedisKey) WithPrefix(prefix string) RedisKey {
	return RedisKey(prefix + ":" + k.String())
}

func (k RedisKey) WithSuffix(suffix string) RedisKey {
	return RedisKey(k.String() + ":" + suffix)
}

// WithKeyValue returns a new RedisKey with the given key-value pairs.
// For example, WithKeyValue("key1", 1, "key2", 2) returns "key:key1:key2:2".
func (k RedisKey) WithKeyValue(pairs ...interface{}) RedisKey {
	if len(pairs)%2 != 0 {
		panic("WithKeyValue requires an even number of arguments")
	}

	var strPairs []string
	for i := 0; i < len(pairs); i += 2 {
		key, okKey := pairs[i].(string)
		value := pairs[i+1]

		if !okKey {
			panic(fmt.Sprintf("WithKeyValue requires pairs of string and interface{}, got %v and %v", pairs[i], pairs[i+1]))
		}

		strValue := fmt.Sprintf("%v", value)
		strPairs = append(strPairs, key+":"+strValue)
	}

	return RedisKey(k.String() + ":" + strings.Join(strPairs, ":"))
}
