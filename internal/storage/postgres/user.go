package postgres

import (
	"go.uber.org/fx"
	"log"
)

type UserStorage struct {
	db *Storage
}

func NewUserStorage(lc fx.Lifecycle, db *Storage) *UserStorage {
	return &UserStorage{db: db}
}

func (p UserStorage) GetUser(username, password string) bool {
	log.Println(username, password)
	return false
}
