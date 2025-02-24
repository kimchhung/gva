package relay_test

import (
	"backend/internal/relay"
	"testing"
)

func TestSetPrevious_WithoutAfter(t *testing.T) {
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 5, relay.NewPaginateGlobalConfig())

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}

	p = relay.PageInfo{}
	p.SetHasPreviousPage(10, 5, relay.NewPaginateGlobalConfig(relay.DisableCount()))

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}
}

func TestSetPrevious_WithAfter(t *testing.T) {
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				After: &after,
			}),
		),
	)

	if !p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasPreviousPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				After: &after,
			}),
			relay.DisableCount(),
		),
	)

	if !p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be true")
	}
}

func TestSetPrevious_WithAfter_TotalCountEqualsEdgesLen(t *testing.T) {
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasPreviousPage(10, 10,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				After: &after,
			}),
		),
	)

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}

	p = relay.PageInfo{}
	p.SetHasPreviousPage(10, 10,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				After: &after,
			}),
			relay.DisableCount(),
		),
	)

	if p.HasPreviousPage {
		t.Errorf("Expected HasPreviousPage to be false")
	}
}

func TestSetNextPage_WithBefore(t *testing.T) {
	before := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				Before: &before,
			}),
		),
	)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				Before: &before,
			}),
			relay.DisableCount(),
		),
	)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_TotalCountEqualsEdgesLen(t *testing.T) {
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 10,
		relay.NewPaginateGlobalConfig(),
	)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 10,
		relay.NewPaginateGlobalConfig(
			relay.DisableCount(),
		),
	)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_First_GreaterThan_EdgesLen(t *testing.T) {
	first := 10
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{First: &first}),
		),
	)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{First: &first}),
			relay.DisableCount(),
		),
	)

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_Last_GreaterThan_EdgesLen(t *testing.T) {
	last := 10
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(
			&relay.PaginateConfig{Last: &last},
		),
	))

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(
			&relay.PaginateConfig{Last: &last},
		),
		relay.DisableCount(),
	))

	if p.HasNextPage {
		t.Errorf("Expected HasNextPage to be false")
	}
}

func TestSetNextPage_After_IsNil_First_IsNil_TotalCount_GreaterThan_First(t *testing.T) {
	first := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				First: &first,
			}),
		),
	)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5,
		relay.NewPaginateGlobalConfig(
			relay.Cursor(&relay.PaginateConfig{
				First: &first,
			}),
			relay.DisableCount(),
		),
	)

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_After_IsNil_Last_IsNil_TotalCount_GreaterThan_Last(t *testing.T) {
	last := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{Last: &last}),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{Last: &last}),
		relay.DisableCount(),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_First_Equals_EdgesLen(t *testing.T) {
	first := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{First: &first}),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{First: &first}),
		relay.DisableCount(),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_First_Equals_EdgesLen_With_After(t *testing.T) {
	first := 5
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{
			First: &first, After: &after,
		}),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{
			First: &first, After: &after,
		}),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_Last_Equals_EdgesLen(t *testing.T) {
	last := 5
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{Last: &last}),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}

	p = relay.PageInfo{}
	p.SetHasNextPage(10, 5, relay.NewPaginateGlobalConfig(
		relay.Cursor(&relay.PaginateConfig{Last: &last}),
		relay.DisableCount(),
	))

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}

func TestSetNextPage_Last_Equals_EdgesLen_With_After(t *testing.T) {
	last := 5
	after := "eyJpZCI6MX0="
	p := relay.PageInfo{}
	p.SetHasNextPage(10, 5, &relay.PaginateGlobalConfig{Last: &last, After: &after})

	if !p.HasNextPage {
		t.Errorf("Expected HasNextPage to be true")
	}
}
