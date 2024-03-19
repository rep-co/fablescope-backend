package database

import "github.com/rep-co/fablescope-backend/wardrobe-auth/data"

type Storage interface {
	CreateAccount(*data.Account) error
	GetAccount(email, password string) (*data.Account, error)
	UpdateAccount(email, password string) (*data.Account, error)
	DeleteAccount(email, password string) error
}
