package types

type CursorDto struct {
	First  *int    `json:"first,omitempty" query:"first"`
	Last   *int    `json:"last,omitempty" query:"last"`
	After  *string `json:"after,omitempty" query:"after"`
	Before *string `json:"before,omitempty" query:"before"`

	// sort="createdAt,-amount,+title",
	// eq order by created_at asc, ammount asc, title asc
	Sorts string `json:"sorts,omitempty" query:"sorts"`

	// whether return count record or not, default false
	IsCount bool `json:"isCount,omitempty" query:"isCount"`

	// ?[id][eq]=1&[id][createdAt][eq]=2024-05-17T09:49:15.466Z
	Filters map[string]any
}
