package database

import "fmt"

var (
	NoResultError       = storageError{Message: "no result error. Record doesn't exist"}
	TransactionError    = storageError{Message: "transaction error. Transaction rolled back"}
	ExecutionError      = storageError{Message: "operation error. Operation can't be executed"}
	RequestTimeoutError = storageError{Message: "request timeout"}
)

type storageError struct {
	Message string
}

func (e *storageError) Error() string {
	return fmt.Sprintf("storage error: %s", e.Message)
}
