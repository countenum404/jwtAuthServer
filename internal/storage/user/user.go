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
	GetUserEmailById(guid string) (string, error)
}

type DefaultUserStorage struct {
	db *postgres.Storage
}

func (p DefaultUserStorage) GetUserEmailById(guid string) (string, error) {
	query := "SELECT email FROM users WHERE id=$1;"
	var email string
	err := p.db.Db.QueryRow(query, guid).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, err
}

func NewDefaultUserStorage(lc fx.Lifecycle, db *postgres.Storage) *DefaultUserStorage {
	return &DefaultUserStorage{db: db}
}

func (p DefaultUserStorage) GetUser(username, password string) bool {
	log.Println(username, password)
	return false
}
