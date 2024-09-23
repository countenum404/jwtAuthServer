package storage

import (
	"go.uber.org/fx"
	"jwtAuth/internal/storage/postgres"
)

var Module = fx.Module("storage", fx.Provide(fx.Annotate(postgres.NewUserStorage, fx.As(new(UserStorage)))))

type UserStorage interface {
	GetUser(username, password string) bool
}
