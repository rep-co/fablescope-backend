package database

import (
	"context"
	"path"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost int = bcrypt.DefaultCost // 10
)

type YDBStorage struct {
	db *ydb.Driver
}

func NewYDBStorage(
	ctx context.Context,
	connString, token string,
) (*YDBStorage, error) {

	db, err := ydb.Open(ctx, connString, ydb.WithAccessTokenCredentials(token))
	if err != nil {
		return nil, err
	}

	return &YDBStorage{
		db: db,
	}, nil
}

func (s *YDBStorage) Init(ctx context.Context) error {
	if err := s.createAccountTable(ctx); err != nil {
		return err
	}
	return nil
}

func (s *YDBStorage) Close(ctx context.Context) error {
	return s.db.Close(ctx)
}

func (s *YDBStorage) createAccountTable(ctx context.Context) error {
	err := s.db.Table().Do(ctx,
		func(ctx context.Context, session table.Session) (err error) {
			return session.CreateTable(ctx, path.Join(s.db.Name(), "account"),
				options.WithColumn("account_id", types.TypeString),
				options.WithColumn("name", types.TypeString),
				options.WithColumn("email", types.TypeString),
				options.WithColumn("password", types.TypeString),
				options.WithPrimaryKeyColumn("account_id"),
			)
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *YDBStorage) CreateAccount(ctx context.Context, account *data.Account) error {
	query := `
        DECLARE $account_id AS String;
        DECLARE $name AS String;
        DECLARE $email AS String;
        DECLARE $password AS String;
        UPSERT INTO account (account_id, name, email, password)
        VALUES ($account_id, $name, $email, $password);
    `

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(account.Password), cost)
	if hashErr != nil {
		return hashErr
	}

	err := s.db.Table().DoTx(ctx,
		func(ctx context.Context, tx table.TransactionActor) error {
			res, err := tx.Execute(ctx,
				query,
				table.NewQueryParameters(
					table.ValueParam("$account_id", types.BytesValue(account.ID[:])),
					table.ValueParam("$name", types.BytesValue([]byte(account.Name))),
					table.ValueParam("$email", types.BytesValue([]byte(account.Email))),
					table.ValueParam("$password", types.BytesValue([]byte(hashedPassword))),
				),
			)
			if err != nil {
				return err
			}
			if err = res.Err(); err != nil {
				return err
			}
			return res.Close()
		}, table.WithIdempotent(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *YDBStorage) GetAccount(
	ctx context.Context,
	email,
	password string,
) (*data.Account, error) {
	return nil, nil
}

func (s *YDBStorage) UpdateAccount(
	ctx context.Context,
	email,
	password string,
) (*data.Account, error) {
	return nil, nil
}

func (s *YDBStorage) DeleteAccount(
	ctx context.Context,
	email,
	password string,
) error {
	return nil
}
