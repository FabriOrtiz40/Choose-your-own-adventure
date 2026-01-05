package story

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Story map[string]Chapter

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}

type handler struct {
	s Story
	t *template.Template
}

type HandlerOption func(*handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		if t != nil {
			h.t = t
		}
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s: s, t: tpl}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("Template error: %v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

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
