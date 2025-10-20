package req

import (
	"net/http"
	"strings"

	_req "github.com/imroc/req/v3"
)

type Client = _req.Client
type Req = _req.Request
type Option func(request *Req)

func Request(url string, options ...Option) *Req {
	r := _req.
		C().
		DisableAutoDecode().
		EnableAutoDecompress().
		NewRequest()

	r.Method = http.MethodGet
	r.SetURL(url)

	for _, option := range options {
		option(r)
	}

	return r
}

func New() *Client {
	return _req.NewClient()
}

func RequestStruct(url string, out any, options ...Option) (code int, body []byte, errs []error) {
	request := Request(url, options...)
	response := request.Do()
	if response.Err != nil {
		return response.GetStatusCode(), response.Bytes(), []error{response.Err}
	}
	response.Into(out)
	return response.GetStatusCode(), response.Bytes(), nil
}

func RequestBytes(url string, options ...Option) (code int, body []byte, errs []error) {
	request := Request(url, options...)
	response := request.Do()
	if response.Err != nil {
		return response.GetStatusCode(), response.Bytes(), []error{response.Err}
	}
	return response.GetStatusCode(), response.Bytes(), nil
}

func GetFullUrl(r *Req) string {
	baseURL := r.RawURL
	query := r.QueryParams.Encode()

	if query == "" {
		return baseURL
	}

	if strings.Contains(baseURL, "?") {
		return baseURL + "&" + query
	}

	return baseURL + "?" + query
}
