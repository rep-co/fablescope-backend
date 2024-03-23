package middlewares

import (
	"context"
	"fmt"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
)

var (
	contextKeyAccountRequest = contextKey("accountRequest")
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares context key " + string(c)
}

func GetAccountRequestKey(ctx context.Context) (*data.AccountRequest, error) {
	if v := ctx.Value(contextKeyAccountRequest); v != nil {
		if v, ok := v.(*data.AccountRequest); ok {
			return v, nil
		}
		err := &KeyHasWrongTypeError{keyName: string(contextKeyAccountRequest)}
		return nil, err
	}
	err := &KeyWasNotFoundError{keyName: string(contextKeyAccountRequest)}
	return nil, err
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
