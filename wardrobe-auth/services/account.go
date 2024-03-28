package services

import (
	"context"
	"time"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/database"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashingCost = bcrypt.DefaultCost // 10
	ydbRequestTTL       = time.Second * 5
)

type AccountService struct {
	accountStorage database.AccountStorage
}

func NewAccountService(
	accountStorage database.AccountStorage,
) *AccountService {
	return &AccountService{
		accountStorage: accountStorage,
	}
}

// Tries to create new account from given request.
//
// Returns database.RequestTimeoutError if transaction took more than 5 sec.
//
// Returns database.TransactionError if account with given email already exists
func (as *AccountService) CreateNewAccount(
	ctx context.Context,
	request *data.AccountRequest,
) error {
	account := data.NewAccount(request)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(account.Password),
		passwordHashingCost,
	)
	if err != nil {
		return err
	}
	account.Password = string(hashedPassword)

	ctxTxDeadline, cancel := context.WithDeadline(
		ctx,
		time.Now().Add(ydbRequestTTL),
	)
	defer cancel()

	err = as.accountStorage.CreateAccount(ctxTxDeadline, account)
	if err != nil {
		return err
	}

	return nil
}

// Tries to authorize given account.
//
// Returns nil, database.RequestTimeoutError if transaction took more than 5 sec.
//
// Returns nil, database.NoResultError if account wasn't found.
//
// Returns nil, bcrypt.ErrMismatchedHashAndPassword if password is wrong
func (as *AccountService) AuthorizeAccount(
	ctx context.Context,
	request *data.AccountRequest,
) (*data.Account, error) {
	ctxTxDeadline, cancel := context.WithDeadline(
		ctx,
		time.Now().Add(ydbRequestTTL),
	)
	defer cancel()

	account, err := as.accountStorage.GetAccount(ctxTxDeadline, request.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(account.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return nil, err
	}

	return account, nil
}
