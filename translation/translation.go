package translation

import (
	"fmt"
	"strings"
)

func Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)
	fmt.Println("Translate: word", word)
	if word != "hello" {
		return ""
	}
	fmt.Println("Translate: language", language)
	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
