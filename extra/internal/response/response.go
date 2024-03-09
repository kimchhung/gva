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

func defaultReponseBase() ReponseBase {
	return ReponseBase{
		Code:       successCode,
		Message:    successMessage,
		httpStatus: successHttpCode,
	}
}

func New(opt ReponseOption, opts ...ReponseOption) *Response {
	resp := &Response{
		ReponseBase: defaultReponseBase(),
	}

	opt(resp)

	for _, op := range opts {
		op(resp)
	}

	return resp
}

type (
	ReponseOption func(resp *Response)

	// ResponseBase represents the basic structure of a response, including
	// a status code, a message.
	ReponseBase struct {
		// expose
		Code    int    `json:"code"`
		Message string `json:"message"`

		// internal use
		httpStatus     int
		lockHttpStatus *int
	}

	Response struct {
		ReponseBase

		// The Data field contains the actual response data
		Data any `json:"data,omitempty"`

		// Meta provides additional information about the data, such as its type or kind.y.
		Meta any `json:"meta,omitempty"`
	}
)

func (r *Response) Parse(c *fiber.Ctx) error {
	if r.lockHttpStatus != nil {
		return c.Status(*r.lockHttpStatus).JSON(r)
	}

	return c.Status(r.httpStatus).JSON(r)
}
