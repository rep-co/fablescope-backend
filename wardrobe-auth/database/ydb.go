package database

import (
	"context"
	"log"
	"path"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

type YDBStorage struct {
	db     *ydb.Driver
	hasher any
}

func NewYDBStorage() (*YDBStorage, error) {
	ctx := context.TODO()
	dsn := "grpcs://ydb.serverless.yandexcloud.net:2135/?database=/ru-central1/b1gih93q5tulltkd8r47/etnav8lc4tnqftk3fu2m"
	token := "t1.9euelZqYyJCQyprOxs-Tz5rMm5SLi-3rnpWaiZCLzM-eipuXkZvPy5mOx5Hl8_ckURhQ-e8oNkRN_N3z92R_FVD57yg2RE38zef1656Vms2WkpOPxs2ej8-PzJqYyp6e7_zF656Vms2WkpOPxs2ej8-PzJqYyp6e.1O-WywPerb4lV0m2stWyFp23nElbqbHnUMtTRbILTyw9-sqQIYThl5I_DuqTZcD8OzMKC-TsNPCYjNeBHM8jCA"

	db, err := ydb.Open(ctx, dsn, ydb.WithAccessTokenCredentials(token))
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

	// TODO: Do hashing
	hashedPassword := account.Password

	log.Printf("%v", account)

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
