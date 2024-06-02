// Package pagi provides utilities for implementing pagination in Go applications.
// It includes structures and methods for calculating page numbers, limits, and offsets
// based on total item counts and user-defined page sizes. This package aims to simplify
// the process of adding pagination to data-driven applications, ensuring efficient
// data retrieval and presentation.
package pagi

type Meta struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Total  *int `json:"total"`
}

type Response[T any] struct {
	Data []T   `json:"data"`
	Meta *Meta `json:"meta"`
}

type Param struct {
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Filter map[string]any `json:"filter"`
	Sort   []string       `json:"sort"`
}
