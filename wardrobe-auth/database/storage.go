package database

import "github.com/rep-co/fablescope-backend/wardrobe-auth/data"

type Storage interface {
	CreateAccount(*data.AccountCredentials) error
	GetAccount(*data.AccountCredentials) (*data.Account, error)
	UpdateAccount(*data.AccountCredentials) (*data.Account, error)
	DeleteAccount(*data.AccountCredentials) error
}
