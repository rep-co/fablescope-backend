package database

import "fmt"

var (
	ExecutionError = storageError{Message: "execution error"}
	NoResultError  = storageError{Message: "no data were recived"}
)

type storageError struct {
	Message string
}

func (e *storageError) Error() string {
	return fmt.Sprintf("Storage error: %s", e.Message)
}
