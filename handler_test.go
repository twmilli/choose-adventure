package cyoa

import (
	"fmt"
	"testing"
)

func TestParseStoryChapterJSON(t *testing.T) {
	parsed := parseStoryChapterJSON("gopher.json")
	fmt.Println(parsed)
	expectedTitle := "The Little Blue Gopher"
	if parsed["intro"].Title != expectedTitle {
		t.Errorf(
			"Expected intro title: %v"+
				"actual: %v",
			expectedTitle, parsed["intro"].Title)
	}

	expectedOptionOneText := "This is getting too weird for me. Let's bail and head back home."
	if parsed["new-york"].Options[0].Text != expectedOptionOneText {
		t.Errorf(
			"Expected option 1 text: %v"+
				"actual: %v",
			expectedTitle, parsed["new-york"].Options[0].Text)
	}
}
