package cyoa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type storyChapter struct {
	Title     string   `json:"title"`
	StoryText []string `json:"story"`
	Options   []Option `json:"options"`
}

// MainHandler Main HTTP Handler interface
type MainHandler struct {
	storyChapters map[string]storyChapter
}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
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
	// fmt.Println(parsed["intro"])
	// return nil
}
