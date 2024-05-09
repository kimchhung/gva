package request_test

import (
	"testing"

	"github.com/kimchhung/gva/extra/internal/request"
)

func TestInit(t *testing.T) {
	t.Run("test parse multiple", func(t *testing.T) {
		q := "filter[money][gt]=100&limit=100"
		e := `{"filter":{"money":{"gt":100}},"limit":100}`

		v := request.ParseUrlQuery(q)

		if e != v {
			t.Errorf("invalid: %v, expected: %v", v, e)
		}
	})

	t.Run("test parse single", func(t *testing.T) {
		q := "filter[money][gt]=100"
		e := `{"filter":{"money":{"gt":100}}}`

		v := request.ParseUrlQuery(q)

		if e != v {
			t.Errorf("invalid: %v, expected: %v", v, e)
		}
	})

	t.Run("test parse bool and number", func(t *testing.T) {
		q := "filter[money][gt]=true&limit=10"
		e := `{"filter":{"money":{"gt":true}},"limit":10}`

		v := request.ParseUrlQuery(q)

		if e != v {
			t.Errorf("invalid: %v, expected: %v", v, e)
		}
	})
}
