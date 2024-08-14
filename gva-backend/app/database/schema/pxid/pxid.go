package pxid

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

// ID implements a xid - a prefixed xid.
type ID string

// newULID returns a new ULID for time.Now() using the default entropy source.
func newId(prefix string) ID {
	xid.New()
	return ID(prefix + seperator + xid.New().String())
}

// New returns a new xid for time.Now() given a prefix. This uses the default entropy source.
func New(prefix string) ID {
	return newId(prefix)
}

// -1 if not found
func (u ID) PrefixIndex() int {
	return strings.Index(string(u), seperator)
}

func (u ID) XID() (xid.ID, error) {
	prefixIndex := u.PrefixIndex()
	if prefixIndex < 0 {
		return xid.ID{}, nil
	}
	xid, err := xid.FromString(string(u)[prefixIndex:])
	if err != nil {
		return xid, err
	}
	return xid, nil
}

func (u ID) String() string {
	return string(u)
}

func (u ID) Prefix() string {
	prefixIndex := u.PrefixIndex()
	if prefixIndex < 0 {
		return ""
	}

	return string(u)[:prefixIndex]
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
	case []uint8:
		*u = ID(src)
	default:
		return fmt.Errorf("xid: unexpected type, %T", src)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}
