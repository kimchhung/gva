package mangatoapi

type Genre struct {
	ID        string
	GenreName string
	Mangas    []Manga
}

type Page struct {
	ID       string
	ImageURL string
}

type Chapter struct {
	ID          string
	MangaID     string
	ChapterName string
	Views       string
	Uploaded    string
	Pages       []Page
}

type Author struct {
	ID     string
	Name   string
	Mangas []Manga
}

type Manga struct {
	ID           string
	Thunmnail    string
	Name         string
	Alternatives string
	Author       Author
	Status       string
	Updated      string
	Views        string
	Rating       string
	Description  string
	Genres       []Genre
	Chapters     []Chapter
}
