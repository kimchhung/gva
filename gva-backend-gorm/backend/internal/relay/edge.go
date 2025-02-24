package relay

import rcursor "backend/internal/relay/cursor"

func convertToEdge[T any](encoder rcursor.Encoder, rows []*T, fields []string, primaryKey string, isConvertAll bool) ([]*Edge[T], error) {
	edges := make([]*Edge[T], len(rows))

	for i, row := range rows {
		var (
			cursor string
			err    error
		)

		if isConvertAll {
			cursor, err = rcursor.Create(encoder, row, fields, primaryKey)
		} else {
			if i == 0 || i == len(rows)-1 {
				cursor, err = rcursor.Create(encoder, row, fields, primaryKey)
			}
		}

		if err != nil {
			return nil, err
		}

		edges[i] = &Edge[T]{
			Cursor: cursor,
			Node:   row,
		}
	}

	return edges, nil
}
