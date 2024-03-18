package data

import "github.com/google/uuid"

type AccountCredentials struct {
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
