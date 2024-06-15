package json

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"slices"

	"errors"
	"fmt"
)

type JSON json.RawMessage

// Convert Json string to bytes,
func NewJSON(value any) (JSON, error) {
	bytes, err := json.Marshal(value)
	return JSON(bytes), err
}

func MustJSON(value any) JSON {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("failed to marshal: %e", err))
	}

	return JSON(bytes)
}

func NewEmptyObject() JSON {
	return JSON([]byte("{}"))
}

func NewJSONNull() JSON {
	return JSON([]byte("null"))
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// Convert Struct to Bytes
func (j *JSON) Set(value interface{}) error {
	bytes, err := json.Marshal(value)
	*j = JSON(bytes)
	return err
}

// Convert Bytes
func (j JSON) Out(out interface{}) error {
	return json.Unmarshal(j, &out)
}

func (j JSON) String() string {
	return string(j)
}

func (j JSON) ArrayInt() []int {
	res := make([]int, 0)
	json.Unmarshal(j, &res)
	return res
}

func (j JSON) Object() map[string]any {
	res := make(map[string]any, 0)
	json.Unmarshal(j, &res)
	return res
}

// Unmarshal | string, remove json whitespace
func (j JSON) Dump() any {
	obj := j.Object()
	if len(obj) == 0 {
		return j.String()
	}
	return obj
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j JSON) IsValid() bool {
	return json.Valid(j)
}

func (j JSON) IsEqual(dest JSON) bool {
	return j.String() == dest.String()
}

func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func IsDeepEqual[Old comparable, New comparable](old Old, new New, omitJsonKeys ...string) bool {
	oldMap := MustJSON(old).Object()
	newMap := MustJSON(new).Object()

	for field := range oldMap {
		if slices.Contains(omitJsonKeys, field) {
			continue
		}

		if !MustJSON(oldMap[field]).IsEqual(MustJSON(newMap[field])) {
			return false
		}
	}

	return true
}
