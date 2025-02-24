package utils

// Get last Element
func Last[E any](items []E) *E {
	length := len(items)
	if length == 0 {
		return nil
	} else {
		return &items[length-1]
	}
}
