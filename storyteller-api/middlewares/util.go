package middlewares

import (
	"bytes"
	"context"
	"github.com/rep-co/fablescope-backend/storyteller-api/data"
)

var (
	contextKeyTags  = contextKey("tags")
	contextKeyStory = contextKey("story")
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares context key " + string(c)
}

func GetTagsKey(ctx context.Context) ([]data.TagName, error) {
	if v := ctx.Value(contextKeyTags); v != nil {
		return v.([]data.TagName), nil
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyTags)}
	return nil, err
}

func GetStoryKey(ctx context.Context) (data.Story, error) {
	if v := ctx.Value(contextKeyStory); v != nil {
		return v.(data.Story), nil
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyStory)}
	return *data.NewStory(""), err
}

type KeyWasNotFoundError struct {
	keyName string
}

// TODO: mb it's better to use strings.Builder
// But anyway the result will be the same
func (m *KeyWasNotFoundError) Error() string {
	var b bytes.Buffer

	b.WriteString("Key was not found: ")
	b.WriteString(m.keyName)

	return b.String()
}
