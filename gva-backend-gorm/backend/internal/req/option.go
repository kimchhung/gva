package req

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// DONE
func SetCurl(curlCmd *string) Option {
	return func(request *Req) {
		method := request.Method
		url := GetFullUrl(request)
		body := request.Body

		// Start the cURL command with the method and URL
		*curlCmd = fmt.Sprintf("curl -X %s '%s'", method, url)

		for key, value := range request.Headers {
			safeKey := key
			safeValue := strings.Join(value, ",")
			*curlCmd += fmt.Sprintf(" -H '%s: %s'", safeKey, safeValue)
		}

		// Add body to the cURL command if it's a POST or PUT request
		if method == http.MethodPost || method == http.MethodPut {
			// Ensure the body is a valid UTF-8 string
			if body != nil {
				safeBody := strings.ReplaceAll(string(body), "'", "\\'")
				*curlCmd += fmt.Sprintf(" -d '%s'", safeBody)
			}
		}
	}
}

// Done
func WithHeader(headers map[string]string) Option {
	return func(request *Req) {
		request.SetHeaders(headers)
	}
}

func WithTimeout(duration time.Duration) Option {
	return func(request *Req) {
		request.GetClient().SetTimeout(duration)
	}
}

// Done
func WithBody(body interface{}) Option {
	return func(request *Req) {
		if body == nil {
			return
		}

		switch body := body.(type) {
		case string:
			request.SetBodyString(body)
		case []byte:
			request.SetBodyBytes(body)
		default:
			request.SetBodyJsonMarshal(body)
		}
	}
}

func WithClient(c *Client) Option {
	return func(request *Req) {
		request.SetClient(c)
	}
}

func WithContext(ctx context.Context) Option {
	return func(request *Req) {
		request.SetContext(ctx)
	}
}

func WithQueries(queries map[string]string) Option {
	return func(request *Req) {
		if len(queries) == 0 {
			return
		}
		request.SetQueryParams(queries)
	}
}

// DONE
// http.MethodPost,http.MethodGet
func WithMethod(method string) Option {
	return func(request *Req) {
		if method != "" {
			request.Method = method
		}
	}
}

// Done
func WithUserAgent(userAgent string) Option {
	return func(request *Req) {
		request.SetHeader("User-Agent", userAgent)
	}
}

func Options(opts ...Option) Option {
	return func(request *Req) {
		for _, opt := range opts {
			opt(request)
		}
	}
}
