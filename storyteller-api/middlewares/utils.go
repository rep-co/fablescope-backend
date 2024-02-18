package middlewares

import (
	"bytes"
	"context"
	"fmt"
)

var (
	contextKeyTags  = contextKey("tags")
	contextKeyStory = contextKey("story")
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares context key " + string(c)
}

func GetTagsKey(ctx context.Context) ([]string, error) {
	if v := ctx.Value(contextKeyTags).([]string); v != nil {
		fmt.Println(v)
		return v, nil
	}
	err := &KeyWasNotFoundError{}
	return nil, err
}

func GetStoryKey(ctx context.Context) (string, error) {
	if v := ctx.Value(contextKeyStory).(string); v != "" {
		fmt.Println(v)
		return v, nil
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyTags)}
	return "", err
}

type KeyWasNotFoundError struct {
	keyName string
}

func (m *KeyWasNotFoundError) Error() string {
	var b bytes.Buffer

	b.WriteString("Key was not found: ")
	b.WriteString(m.keyName)

	return b.String()
}
