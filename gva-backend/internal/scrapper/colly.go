package scrapper

import (
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/gocolly/colly/v2/extensions"
)

func WriteHtmlFile(c *colly.Collector) {
	c.OnResponse(func(r *colly.Response) {
		// Print the response status code
		log.Println("Response status code:", r.StatusCode)

		// Extract the HTML content of the page
		html := string(r.Body)

		filename := strings.ReplaceAll(strings.TrimPrefix(r.Request.URL.String(), "/https://"), "/", "__")
		// Write the HTML content to a local file named url.html
		err := os.WriteFile(filename+".html", []byte(html), 0644)
		if err != nil {
			log.Println("Error writing HTML to file:", err)
			return
		}

		log.Println("HTML written to file url.html")
	})
}

var Debug = true

func NewCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
		colly.AllowURLRevisit(),
	)

	if Debug {
		WriteHtmlFile(c)
	}

	extensions.RandomUserAgent(c)
	return c
}
