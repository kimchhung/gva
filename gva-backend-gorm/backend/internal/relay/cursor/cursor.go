package cursor

import (
	"fmt"
	"backend/internal/relay/relayt"
	"backend/internal/relay/utils"
	"reflect"
	"strings"
)

type Encoder interface {
	Encode(v map[string]any) (string, error)
	Decode(raw string) (map[string]any, error)
}

func Create[T any](encoder Encoder, row *T, fields []string, primaryKey string) (string, error) {
	reflectRow := reflect.ValueOf(*row)

	if len(fields) == 0 {
		_, key := utils.SplitOrderKey(primaryKey)
		cursorMap := map[string]any{
			primaryKey: reflectRow.FieldByNameFunc(fieldWithColumnIsEqual(key)).Interface(),
		}

		return encoder.Encode(cursorMap)
	}

	cursorMap := map[string]any{}
	for _, key := range fields {
		_, field := utils.SplitOrderKey(key)
		cursorMap[field] = reflectRow.FieldByNameFunc(fieldWithColumnIsEqual(field)).Interface()
	}

	return encoder.Encode(cursorMap)
}

func fieldWithColumnIsEqual(field string) func(key string) bool {
	return func(key string) bool {
		return strings.ToLower(key) == strings.ReplaceAll(field, "_", "")
	}
}

func ParseCursor(encoder Encoder, dataRaw string) (map[string]any, error) {
	cursor, err := encoder.Decode(dataRaw)
	if err != nil {
		return nil, relayt.NewErrRelay(fmt.Sprintf("invalid cursor format %v", err))
	}

	return cursor, nil
}
