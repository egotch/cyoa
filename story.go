// functions and types for creating a Story struct from parsed json
package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// JsonStory - returns Story struct
// parses json data file into a Story struct
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)

	if err != nil {
		return nil, err
	} else {
		return story, nil
	}

}
