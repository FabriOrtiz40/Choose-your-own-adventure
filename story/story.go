package story

import (
	"io"
	"encoding/json"
)

type Story map [string]Chapter

func JsonStory(r io.Reader) (Story, error) {
    var story Story
    d := json.NewDecoder(r)
    if err := d.Decode(&story); err != nil {
        return nil, err
    }
    return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

/*type Demo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}*/
