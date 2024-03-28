package data

import "github.com/google/uuid"

type AccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountResponse struct {
	ID string `json:"account_id"`
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
