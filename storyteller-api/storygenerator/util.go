package storygenerator

import (
	"strings"

	"github.com/rep-co/fablescope-backend/storyteller-api/data"
)

func tagNamesToString(tagNames []data.TagName) string {
	var sb strings.Builder
	for i, tag := range tagNames {
		sb.WriteString(tag.Name)
		if i < len(tagNames)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
