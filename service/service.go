package service

import (
	"sfdbhw/storage"
)

type Services struct {
	taskStorage storage.Storages
}

func NewServices(storage storage.Storages) Services {
	return Services{
		taskStorage: storage,
	}
}
