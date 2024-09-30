package cyoa

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
