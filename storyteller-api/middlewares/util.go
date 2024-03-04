package middlewares

import (
	"context"
	"fmt"

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
		if v, ok := v.([]data.TagName); ok {
			return v, nil
		}
		err := &KeyHasWrongTypeError{keyName: string(contextKeyTags)}
		return nil, err
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyTags)}
	return nil, err
}

func GetStoryKey(ctx context.Context) (*data.Story, error) {
	if v := ctx.Value(contextKeyStory); v != nil {
		if v, ok := v.(*data.Story); ok {
			return v, nil
		}
		err := &KeyHasWrongTypeError{keyName: string(contextKeyStory)}
		return data.NewStoryEmpty(), err
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyStory)}
	return data.NewStoryEmpty(), err
}

type KeyWasNotFoundError struct {
	keyName string
}

func (m *KeyWasNotFoundError) Error() string {
	return fmt.Sprintf("key was not found: %s", m.keyName)
}

type KeyHasWrongTypeError struct {
	keyName string
}

func (m *KeyHasWrongTypeError) Error() string {
	return fmt.Sprintf("key has wrong type: %s. Can't get", m.keyName)
}
