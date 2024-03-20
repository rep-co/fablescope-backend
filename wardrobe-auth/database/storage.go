package database

import (
	"context"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
)

type Storage interface {
	CreateAccount(
		ctx context.Context,
		account *data.Account,
	) error

	GetAccount(
		ctx context.Context,
		email,
		password string,
	) (*data.Account, error)

	UpdateAccount(
		ctx context.Context,
		email,
		password string,
	) (*data.Account, error)

	DeleteAccount(
		ctx context.Context,
		email,
		password string,
	) error
}
