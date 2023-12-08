package database

import (
	"gzfs/Go~Edita/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/libsql/libsql-client-go/libsql"
)

type TursoStorage struct {
	TursoDB *sqlx.DB
}

func New(configEnv config.Env) (*TursoStorage, error) {
	tursoDB, err := sqlx.Connect("libsql", configEnv.DB_URL)
	if err != nil {
		return nil, err
	}

	return &TursoStorage{
		TursoDB: tursoDB,
	}, nil
}

func (tursoStorage *TursoStorage) Close() error {
	return tursoStorage.TursoDB.Close()
}
