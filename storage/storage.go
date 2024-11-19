package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sfdbhw/storage/postgres"
)

type Storages struct {
	Tasker postgres.Tasker
}

func NewStorages(db *pgxpool.Pool) Storages {
	taskStorage := postgres.NewTaskStorage(db)
	return Storages{
		Tasker: &taskStorage,
	}
}
