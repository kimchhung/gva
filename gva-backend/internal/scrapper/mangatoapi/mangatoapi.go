package mangatoapi

import (
	"context"
	"errors"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gva/internal/logging"
	"github.com/gva/internal/scrapper"
	"github.com/gva/internal/validator"
	"github.com/iancoleman/strcase"

	"github.com/rs/zerolog"
)

type ManganatoApi struct {
	colly *colly.Collector
	log   zerolog.Logger
}

func NewManganatoApi(colly *colly.Collector, log zerolog.Logger) *ManganatoApi {
	mlog := log.With().Str("api", "ManganatoApi").Logger()
	m := &ManganatoApi{
		colly: colly,
		log:   mlog,
	}
	return m
}

type SearchDetail struct {
	Domain         string   `json:"domain" validate:"required"`
	LastChapterUrl string   `json:"chapterUrl" validate:"required"`
	MangaId        string   `json:"mangaId" validate:"required"`
	DetailUrl      string   `json:"detailUrl" validate:"required"`
	Title          string   `json:"title" validate:"required"`
	Authors        []string `json:"authors" validate:"required"`
	//Jul 05,2022 - 17:08
	UpdatedAt string `json:"updatedAt" validate:"required"`
	Thumpnail string `json:"thumpnail" validate:"required"`
}

/*
data example:

	{
	    "keyword": "the_evil_ring",
	    "details": [
	        {
	           "chapterUrl": "https://chapmanganato.to/manga-eq981751/chapter-167",
	           "mangaId": "manga-eq981751",
	           "detailUrl": "https://chapmanganato.to/manga-eq981751",
	           "title": "The Devil Ring Chapter 167",
	           "author": "Zhang Sanfeng",
	           "updatedAt": "Updated : Feb 26,2023 - 23:56",
	           "thumpnailtedAt": "https://avt.mkklcdnv6temp.com/26/n/18-1583498117.jpg"
	        },
	    ]
	}
*/
type SearchResult struct {
	Keyword string         `json:"keyword"`
	Details []SearchDetail `json:"details" `
}

func (m *ManganatoApi) SearchManga(ctx context.Context, name string) (*SearchResult, []error) {
	keyword := strcase.ToSnake(name)
	collector := m.colly.Clone()

	url := urlSearchManga(keyword)

	log := zerolog.Ctx(ctx).With().
		Str("url", url).
		Str("keyword", keyword).Logger()

	var errs []error
	searchResult := SearchResult{
		Keyword: keyword,
	}

	collector.OnHTML(".search-story-item", func(h *colly.HTMLElement) {
		detail := SearchDetail{}
		detail.DetailUrl = h.ChildAttr("a.item-img", "href")
		paths := strings.SplitAfter(detail.DetailUrl, "/")

		logging.Log(paths)

		if len(paths) > 0 {
			detail.MangaId = strings.Replace(paths[len(paths)-1], "manga-", "", 1)
			detail.Domain = strings.Join(paths[0:len(paths)-1], "")
		}
		detail.LastChapterUrl = h.ChildAttr(".item-right > a:nth-child(2)", "href")
		detail.Title = h.ChildAttr(".item-right > a:nth-child(2)", "title")

		detail.Authors = strings.Split(h.ChildText(".item-right span.item-author"), ",")
		for i, au := range detail.Authors {
			detail.Authors[i] = strings.TrimSpace(au)
		}

		detail.UpdatedAt = h.ChildText(".item-right span.item-author+span")
		detail.Thumpnail = h.ChildAttr(".search-story-item .img-loading", "src")

		if err := validator.ValidateStruct(detail); err != nil {
			errs = append(errs, err)
			return
		}

		searchResult.Details = append(searchResult.Details, detail)
	})

	collector.OnError(func(r *colly.Response, e error) {
		errs = append(errs, e)
	})

	if err := collector.Visit(url); err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	collector.Wait()
	if isSucess := len(searchResult.Details) > 0; isSucess {
		log.Info().Int("detailCount", len(searchResult.Details)).Msg("success")
		return &searchResult, nil
	}

	errs = append(errs, errors.New("no manga found"))
	return nil, errs
}

type FetchDetailChapter struct {
	Name      string   `json:"name" validate:"required"`
	Url       string   `json:"url" validate:"required"`
	Uploaded  string   `json:"uploaded" validate:"required"`
	ChapterId string   `json:"chapterId" validate:"required"`
	ImgUrls   []string `json:"ImgUrls"`
}

type FetchDetailGenres struct {
	Name    string `json:"name" validate:"required"`
	Url     string `json:"url" validate:"required"`
	GenreId string `json:"genreId" validate:"required"`
}

type FetchMangaDetailResult struct {
	Title   string   `json:"title" validate:"required"`
	Authors []string `json:"authors" validate:"required"`
	//Jul 05,2022 - 17:08
	UpdatedAt   string               `json:"updatedAt" validate:"required"`
	Rating      string               `json:"Rating" validate:"required"`
	Thumpnail   string               `json:"thumpnail" validate:"required"`
	Description string               `json:"description" validate:"required"`
	Status      string               `json:"status" validate:"required"`
	Chapters    []FetchDetailChapter `json:"chapters" validate:"required,dive"`
	Genres      []FetchDetailGenres  `json:"genres" validate:"dive"`
	Alternative []string             `json:"alternative" validate:"required"`
}

func (m *ManganatoApi) FetctMangaDetail(ctx context.Context, domain, mangaId string) (*FetchMangaDetailResult, []error) {
	var (
		detail FetchMangaDetailResult
		errs   []error
	)

	paths := ConstructMangatoPath(domain, mangaId)
	url := paths.Domain + paths.DetailPath
	collector := m.colly.Clone()

	collector.OnHTML("#panel-story-info-description", func(h *colly.HTMLElement) {
		detail.Description = strings.TrimSpace(strings.ReplaceAll(h.Text, "Description :", ""))
	})

	collector.OnHTML(".panel-story-info", func(h *colly.HTMLElement) {
		detail.Thumpnail = h.ChildAttr(".story-info-left img.img-loading", "src")
		detail.Title = h.ChildAttr(".story-info-left img.img-loading", "title")

		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			key := strings.ReplaceAll(h.ChildText(".table-label"), " ", "")
			key = strings.ToLower(strings.ReplaceAll(key, ":", ""))

			switch true {
			case strings.Contains(key, "alternative"):
				detail.Alternative = strings.Split(strings.TrimSpace(h.ChildText(".table-value")), ";")
			case strings.Contains(key, "author"):
				h.ForEach(".table-value a", func(i int, h *colly.HTMLElement) {
					author := h.Text
					if author != "" {
						detail.Authors = append(detail.Authors, author)
					}
				})
			case strings.Contains(key, "status"):
				detail.Status = h.ChildText(".table-value")
			case strings.Contains(key, "genre"):
				detail.Genres = []FetchDetailGenres{}
				h.ForEach(".table-value a", func(i int, h *colly.HTMLElement) {
					genre := FetchDetailGenres{}
					genre.Name = h.Text
					genre.Url = h.Attr("href")
					parts := strings.Split(genre.Url, "-")
					if len(parts) > 0 {
						genre.GenreId = parts[len(parts)-1]
					}
					if genre.Name != "" {
						detail.Genres = append(detail.Genres, genre)
					}
				})
			}
		})

		detail.Alternative = strings.Split(h.ChildAttr(".story-info-left img.img-loading", "title"), ";")
	})

	collector.OnHTML(".story-info-right-extent", func(h *colly.HTMLElement) {
		updated := h.ChildText("p:nth-child(1) .stre-value")
		detail.UpdatedAt = updated
		detail.Rating = h.ChildText("#rate_row_cmd > em > em:nth-child(2) > em > em:nth-child(1)")
	})

	collector.OnHTML(".variations-tableInfo", func(h *colly.HTMLElement) {
		alternatives := h.ChildText("tr:nth-child(1) .table-value")
		status := h.ChildText("tr:nth-child(3) .table-value")

		detail.Alternative = strings.Split(alternatives, ";")
		for i, al := range detail.Alternative {
			detail.Alternative[i] = strings.TrimSpace((al))
		}
		detail.Status = status
	})

	collector.OnHTML(".row-content-chapter li.a-h", func(h *colly.HTMLElement) {
		ch := FetchDetailChapter{}
		ch.Url = h.ChildAttr("a.chapter-name", "href")
		if idx := strings.Index(ch.Url, "chapter-"); idx != -1 {
			ch.ChapterId = ch.Url[idx+1:]
		}
		ch.Name = h.ChildText("a.chapter-name")
		ch.Uploaded = h.ChildAttr("span.chapter-time", "title")
		if strings.TrimSpace(ch.Name) != "" {
			detail.Chapters = append(detail.Chapters, ch)
		}
	})

	scrapper.WriteHtmlFile(collector)
	if err := collector.Visit(url); err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	collector.Wait()
	if err := validator.ValidateStruct(detail); err != nil {
		errs = append(errs, err)
	}

	if isSucess := len(detail.Chapters) > 0; isSucess {
		return &detail, nil
	}

	errs = append(errs, errors.New("no manga detail found"))
	return nil, errs
}
