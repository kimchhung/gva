package mangatoapi

import (
	"fmt"
	"strings"
)

func urlSearchManga(name string) string {
	return fmt.Sprintf("https://manganato.com/search/story/%s", name)
}

type MangatoUrlPath struct {
	Domain        string
	DetailPath    string
	ChapterPathFn func(chapter int) string
}

func ConstructMangatoPath(domain, mangaId string) *MangatoUrlPath {
	path := &MangatoUrlPath{
		Domain:     strings.TrimSuffix(domain, "/"),
		DetailPath: fmt.Sprintf("/manga-%v", mangaId),
		ChapterPathFn: func(chapter int) string {
			return fmt.Sprintf("/manga-%v/chapter-%v", mangaId, chapter)
		},
	}

	return path
}
