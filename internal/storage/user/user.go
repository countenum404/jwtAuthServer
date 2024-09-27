package user

import (
	"go.uber.org/fx"
	"jwtAuth/internal/storage/postgres"
	"log"
)

var Module = fx.Module("UserStorage",
	fx.Provide(fx.Annotate(NewDefaultUserStorage, fx.As(new(Storage)))),
)

type Storage interface {
	GetUser(username, password string) bool
}

type DefaultUserStorage struct {
	db *postgres.Storage
}

func NewDefaultUserStorage(lc fx.Lifecycle, db *postgres.Storage) *DefaultUserStorage {
	return &DefaultUserStorage{db: db}
}

func (p DefaultUserStorage) GetUser(username, password string) bool {
	log.Println(username, password)
	return false
}
