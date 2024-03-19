package database

import (
	"database/sql"

	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	_ "github.com/ydb-platform/ydb-go-sdk/v3"
)

type YDBStorage struct {
	db     *sql.DB
	hasher any
}

func NewYDBStorage() (*YDBStorage, error) {
	connStr := "grpcs://ydb.serverless.yandexcloud.net:2135/ru-central1/b1gih93q5tulltkd8r47/etnav8lc4tnqftk3fu2m?token="
	token := "t1.9euelZrOk5SczomLypOXlImOlImLle3rnpWaiZCLzM-eipuXkZvPy5mOx5Hl8_cTGhxQ-e9nQHhh_N3z91NIGVD572dAeGH8zef1656Vmp6cnpeUiciJkMiQmpGJyIzH7_zF656Vmp6cnpeUiciJkMiQmpGJyIzH.AvLrBtEmKuEtJXONEcEGhNaacbgKwdlnZaC7zxVkN3LgpZzKclsTCPIUCu6uPwn6wos6VVa7wowFTTUyEGl9Cw"

	db, err := sql.Open("ydb", connStr+token)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &YDBStorage{
		db: db,
	}, nil
}

func (s *YDBStorage) Init() error {
	if err := s.createAccountTable(); err != nil {
		return err
	}
	return nil
}

func (s *YDBStorage) createAccountTable() error {
	query :=
		`CREATE TABLE account (
            account_id String,
            name String,
            email String,
            password String,
            PRIMARY KEY (account_id)
        );`

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, execErr := tx.Exec(query)
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *YDBStorage) CreateAccount(account *data.Account) error {
	query :=
		`REPLACE INTO account (account_id, name, email, password) VALUES
            (?, ?, ?, ?);
        `

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// TODO: Hash the password
	_, execErr := tx.Exec(query, account.ID, account.Name, account.Email, account.Password)
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *YDBStorage) GetAccount(email, password string) (*data.Account, error) {
	return nil, nil
}

func (s *YDBStorage) UpdateAccount(email, password string) (*data.Account, error) {
	return nil, nil
}

func (s *YDBStorage) DeleteAccount(email, password string) error {
	return nil
}
