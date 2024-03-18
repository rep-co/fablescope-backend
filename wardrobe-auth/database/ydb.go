package database

import (
	"database/sql"
	_ "github.com/ydb-platform/ydb-go-sdk/v3"
)

type YDBStorage struct {
	db *sql.DB
}

func NewYDBStorage() (*YDBStorage, error) {
	connStr := "grpcs://ydb.serverless.yandexcloud.net:2135/?database=/ru-central1/b1gih93q5tulltkd8r47/etnav8lc4tnqftk3fu2m&token="
	token := "t1.9euelZqUl87Ij8fPzJHJj8-Jnc2VkO3rnpWaiZCLzM-eipuXkZvPy5mOx5Hl9PcfeBxQ-e9ACgSP3fT3XyYaUPnvQAoEj83n9euelZqOlYuXlpfIyp6KjIvPncaWie_8xeuelZqOlYuXlpfIyp6KjIvPncaWiQ.WNt--gfZNjhKNOQFwyNIwuFczO4hSkjQUFAeBVYmSpt9boUWyf2MimzRYK3mSkBeKrFQUU0wp8qnBFeGfJ_mBw"

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

func (s *YDBStorage) CreateAccount(*data.Account) error
