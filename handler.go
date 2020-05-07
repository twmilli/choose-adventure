package cyoa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type storyChapter struct {
	Title     string   `json:"title"`
	StoryText []string `json:"story"`
	Options   []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// MainHandler Main HTTP Handler interface
type MainHandler struct {
	storyChapters map[string]storyChapter
	tmpl          *template.Template
}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	if key == "" {
		key = "intro"
	}
	data, ok := h.storyChapters[key]
	if !ok {
		fmt.Fprintf(w, "Unknown route %v\n", r.URL.Path)
	}
	h.tmpl.Execute(w, data)
}

// NewMainHandler creates a new main handler
func NewMainHandler() MainHandler {
	return MainHandler{
		storyChapters: parseStoryChapterJSON("gopher.json"),
		tmpl:          template.Must(template.ParseFiles("index.html")),
	}
}

func parseStoryChapterJSON(fname string) map[string]storyChapter {
	jsonFile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var parsed map[string]storyChapter
	json.Unmarshal([]byte(byteValue), &parsed)

	return parsed
}
