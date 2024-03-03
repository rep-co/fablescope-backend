package data

import (
	"strings"
	"unicode/utf8"
)

type Story struct {
	Title string `json:"title"`
	Moral string `json:"moral"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Content string `json:"content"`
}

func NewStoryEmpty() *Story {
	return &Story{}
}

func NewStory(rawStory string) *Story {
	rawStory = validateRawStory(rawStory)

	// Ger indexes
	titleIndex := strings.Index(rawStory, "TITLE:")
	moralIndex := strings.Index(rawStory, "MORAL:")
	textIndex := strings.Index(rawStory, "TEXT:")

	// Extract data
	title := rawStory[titleIndex+len("TITLE: ") : moralIndex]
	moral := rawStory[moralIndex+len("MORAL: ") : textIndex]
	text := rawStory[textIndex+len("TEXT: "):]

	// Split into pages
	pages := splitIntoPages(text)

	return &Story{
		Title: title,
		Moral: moral,
		Pages: pages,
	}
}

func validateRawStory(rawStory string) string {
	rawStory = strings.ReplaceAll(rawStory, "\n\n", " ")
	rawStory = strings.ReplaceAll(rawStory, "*", "")
	return rawStory
}

func splitIntoPages(text string) []Page {
	var pageBuilder strings.Builder
	pageSize := 500
	words := strings.Fields(text)
	var pages []Page
	currentLength := 0

	for _, word := range words {
		wordLength := utf8.RuneCountInString(word) + 1
		if currentLength+wordLength > pageSize {
			pages = append(pages, Page{Content: pageBuilder.String()})
			pageBuilder.Reset()
			currentLength = 0
		}
		pageBuilder.WriteString(word)
		pageBuilder.WriteString(" ")
		currentLength += wordLength
	}

	if currentLength > 0 {
		pages = append(pages, Page{Content: pageBuilder.String()})
	}

	return pages
}
