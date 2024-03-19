package database

import "github.com/rep-co/fablescope-backend/wardrobe-auth/data"

type Storage interface {
	CreateAccount(*data.AccountRequest) error
	GetAccount(*data.AccountRequest) (*data.Account, error)
	UpdateAccount(*data.AccountRequest) (*data.Account, error)
	DeleteAccount(*data.AccountRequest) error
}
