package main

import "github.com/rep-co/fablescope-backend/wardrobe-auth/data"

type Storage interface {
	CreateUser(*data.Account) error
	GetAccount(*data.Account) error
	UpdateAccount(*data.Account) error
	DeleteAccount(*data.Account) error
}
