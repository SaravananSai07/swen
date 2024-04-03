package swen

type NewsItem struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Image     string `json:"image"`
	Source    string `json:"source"`
	SourceURL string `json:"sourceUrl"`
}
