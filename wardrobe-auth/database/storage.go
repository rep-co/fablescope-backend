package database

import "github.com/rep-co/fablescope-backend/wardrobe-auth/data"

type Storage interface {
	CreateAccount(*data.Account) error
	GetAccount(*data.Account) error
	UpdateAccount(*data.Account) error
	DeleteAccount(*data.Account) error
}
