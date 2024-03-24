package database

import "fmt"

var (
	ExecutionError   = storageError{Message: "YQL execution error"}
	NoResultError    = storageError{Message: "no result error. Record doesn't exist"}
	TransactionError = storageError{Message: "transaction error. Transaction rolled back"}
)

type storageError struct {
	Message string
}

func (e *storageError) Error() string {
	return fmt.Sprintf("storage error: %s", e.Message)
}
