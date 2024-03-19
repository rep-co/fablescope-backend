package data

import "github.com/google/uuid"

type AccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}

func NewAccount(request *AccountRequest) *Account {
	return &Account{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
