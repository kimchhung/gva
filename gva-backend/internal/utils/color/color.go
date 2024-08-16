package color

import "fmt"

type ColorFunc func(string) string

// NewColorFunc creates a new function that prints its input in the specified color.
func NewColorFunc(colorCode int) ColorFunc {
	return func(s string) string {
		return fmt.Sprintf("\033[%dm%s\033[0m", colorCode, s)
	}
}

var (
	Black   = NewColorFunc(30)
	Red     = NewColorFunc(31)
	Green   = NewColorFunc(32)
	Yellow  = NewColorFunc(33)
	Blue    = NewColorFunc(34)
	Magenta = NewColorFunc(35)
	Cyan    = NewColorFunc(36)
	White   = NewColorFunc(37)
)

func MethodColor(method string) string {
	switch method {
	case "GET":
		return Blue("GET")
	case "POST":
		return Green("POST")
	case "PUT":
		return Yellow("PUT")
	case "DELETE":
		return Red("DELETE")
	case "PATCH":
		return Yellow("PATCH")
	case "HEAD":
		return White("HEAD")
	case "OPTIONS":
		return Blue("OPTIONS")
	default:
		return Black(method)
	}
}
