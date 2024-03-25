package database

import (
	"context"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
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
		switch {
		case ctx.Err() != nil:
			return &RequestTimeoutError
		default:
			fmt.Println(err)
			return &ExecutionError
		}
	}

	return nil
}

func (s *YDBStorage) CreateAccount(
	ctx context.Context,
	account *data.Account,
) error {
	query := `
        DECLARE $account_id AS String;
        DECLARE $name AS String;
        DECLARE $email AS String;
        DECLARE $password AS String;
        UPSERT INTO account (account_id, name, email, password)
        VALUES ($account_id, $name, $email, $password);
    `

	err := s.db.Table().DoTx(ctx,
		func(ctx context.Context, tx table.TransactionActor) error {
			res, err := tx.Execute(ctx,
				query,
				table.NewQueryParameters(
					table.ValueParam("$account_id", types.BytesValue([]byte(account.ID.String()))),
					table.ValueParam("$name", types.BytesValue([]byte(account.Name))),
					table.ValueParam("$email", types.BytesValue([]byte(account.Email))),
					table.ValueParam("$password", types.BytesValue([]byte(account.Password))),
				),
			)
			if err != nil {
				return &TransactionError
			}
			defer res.Close()

			return res.Err()
		}, table.WithIdempotent(),
	)
	if err != nil {
		switch {
		case ctx.Err() != nil:
			return &RequestTimeoutError
		default:
			return err
		}
	}

	return nil
}

func (s *YDBStorage) GetAccount(
	ctx context.Context,
	email string,
) (*data.Account, error) {
	query := `
        DECLARE $email AS String;
        DECLARE $password AS String;
        SELECT account_id, name, email, password
        FROM account
        WHERE email = $email;
    `
	var account data.Account
	var passwordHashString []byte

	err := s.db.Table().DoTx(ctx,
		func(ctx context.Context, tx table.TransactionActor) error {
			res, err := tx.Execute(ctx,
				query,
				table.NewQueryParameters(
					table.ValueParam("$email", types.BytesValue([]byte(email))),
				),
			)
			if err != nil {
				return &TransactionError
			}
			defer res.Close()

			if res.NextResultSet(ctx); res.CurrentResultSet().RowCount() == 0 {
				return &NoResultError
			}

			for res.NextRow() {
				err = res.ScanNamed(
					named.Required("account_id", &passwordHashString),
					named.Required("name", &account.Name),
					named.Required("email", &account.Email),
					named.Required("password", &account.Password),
				)
				if err != nil {
					return fmt.Errorf("account: count not scan account: %w", err)
				}

				account.ID, err = uuid.ParseBytes(passwordHashString)
				if err != nil {
					return fmt.Errorf("account: id is not a uuid format: %w", err)
				}
			}

			return res.Err()
		}, table.WithIdempotent(),
	)
	if err != nil {
		switch {
		case ctx.Err() != nil:
			return nil, &RequestTimeoutError
		default:
			return nil, err
		}
	}

	return &account, nil
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
