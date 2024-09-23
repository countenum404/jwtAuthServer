package service

import (
	"go.uber.org/fx"
	"jwtAuth/internal/storage"
	"log"
)

var Module = fx.Module("service", fx.Provide(NewDefaultUserService, fx.Annotate(NewDefaultUserService, fx.As(new(UserService)))))

type UserService interface {
	GetToken(username, password string) string
}

type DefaultUserService struct {
	storage storage.UserStorage
}

func NewDefaultUserService(lc fx.Lifecycle, storage storage.UserStorage) *DefaultUserService {
	return &DefaultUserService{storage: storage}
}

func (u *DefaultUserService) GetToken(username, password string) string {
	log.Println(username, password)
	return "token"
}
