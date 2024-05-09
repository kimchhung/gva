package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	successCode     = 0
	successMessage  = http.StatusText(http.StatusOK)
	successHttpCode = http.StatusOK
)

func defaultReponseBase() *Response {
	return &Response{
		Code:       successCode,
		Message:    successMessage,
		httpStatus: successHttpCode,
	}
}

func New(opts ...ReponseOption) *Response {
	resp := defaultReponseBase()

	for _, op := range opts {
		op(resp)
	}

	return resp
}

type (
	ReponseOption func(resp *Response)

	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`

		// internal use
		httpStatus int

		// The Data field contains the actual response data
		Data any `json:"data,omitempty"`

		// Meta provides additional information about the data, such as its type or kind.y.
		Meta any `json:"meta,omitempty"`
	}
)

func (r *Response) Parse(c *fiber.Ctx) error {
	return c.Status(r.httpStatus).JSON(r, fiber.MIMEApplicationJSONCharsetUTF8)
}
