// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package xid implements the xid type.
// A xid is an identifier that is a two-byte prefixed ULIDs, with the first two bytes encoding the type of the entity.
package xid

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/xid"
)

const (
	seperator = "_"
)

var _ interface {
	graphql.Marshaler
	graphql.Unmarshaler
} = (*ID)(nil)

// ID implements a xid - a prefixed ULID.
type ID string

// newULID returns a new ULID for time.Now() using the default entropy source.
func newId(prefix string) ID {
	return ID(prefix + seperator + xid.New().String())
}

// MustNew returns a new xid for time.Now() given a prefix. This uses the default entropy source.
func MustNew(prefix string) ID {
	return newId(prefix)
}

// -1 if not found
func (u ID) PrefixIndex() int {
	return strings.Index(string(u), seperator)
}

func (u ID) Prefix() string {
	prefixIndex := u.PrefixIndex()
	if prefixIndex < 0 {
		return ""
	}

	return string(u)[:prefixIndex+1]
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("xid: expected a value")
	}
	switch src := src.(type) {
	case string:
		*u = ID(src)
	case ID:
		*u = src
	default:
		return fmt.Errorf("xid: unexpected type, %T", src)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}
