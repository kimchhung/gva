package req_test

import (
	"backend/internal/req"
	"net/http"
	"testing"
)

// BenchmarkRequestBytes-8   	       1	3001499333 ns/op	    9224 B/op	      73 allocs/op
func BenchmarkRequestBytes(b *testing.B) {
	server := "http://example.com" // Use your test server URL here
	for i := 0; i < b.N; i++ {
		_, _, _ = req.RequestBytes(server)
	}
}

// BenchmarkHTTPRequest-8   	       3	 583241444 ns/op	   18480 B/op	     142 allocs/op
func BenchmarkHTTPRequest(b *testing.B) {
	server := "http://example.com" // Use your test server URL here
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(server)
		if err != nil {
			b.Fatal(err)
		}
		defer resp.Body.Close()
	}
}
