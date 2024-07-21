package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gva/env"
	"github.com/labstack/echo/v4"
)

type httpRespWriter struct {
	mw io.Writer
	http.ResponseWriter
}

func (hrw *httpRespWriter) Write(p []byte) (int, error) {
	return hrw.mw.Write(p)
}

func TraceDebug(cfg *env.Config) echo.MiddlewareFunc {
	skipper := func(c echo.Context) bool {
		if !cfg.IsDev() {
			return true
		}

		for _, bad := range []string{"docs", "playground", "query"} {
			if strings.Contains(c.Request().URL.String(), bad) {
				return true
			}
		}

		return false
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if skipper(c) {
				return next(c)
			}
			startTime := time.Now()
			// Request
			reqBody := []byte{}
			if c.Request().Body != nil { // Read
				reqBody, _ = io.ReadAll(c.Request().Body)
			}
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			// Response
			resBody := new(bytes.Buffer)
			newWriter := &httpRespWriter{
				mw:             io.MultiWriter(c.Response().Writer, resBody),
				ResponseWriter: c.Response().Writer,
			}
			c.Response().Writer = newWriter

			if err = next(c); err != nil {
				c.Error(err)
			}

			// log
			const (
				colorReset = "\033[0m"
				colorBlue  = "\033[34m" // Blue for request details
				colorGreen = "\033[32m" // Green for response body
				colorCyan  = "\033[36m" // Cyan for URL path
			)

			// Get the HTTP method and URL path from the context
			executionTime := time.Since(startTime)
			statusCode := c.Response().Status
			query := c.Request().URL.Query()

			method := c.Request().Method
			path := c.Request().URL.Path

			// Print the request method and URL path in cyan
			fmt.Printf("\n%s[in] -> %s %s:%s\n", colorCyan, method, path, colorReset)

			// Print the request body in blue
			if query.Encode() != "" {
				fmt.Printf("\n%sReq query:%v\n%s\n", colorBlue, colorReset, query)
			}

			if len(reqBody) > 0 {
				fmt.Printf("\n%sReq Body:%s\n%s\n", colorBlue, colorReset, string(reqBody))
			}

			// Print the response body in green
			resp := resBody.Bytes()
			fmt.Printf("\n%s[out] <- %v t=%v respBody:%s\n%s\n", colorGreen, statusCode, executionTime, colorReset, string(resp))
			return
		}
	}
}
