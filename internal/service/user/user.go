package user

import (
	"go.uber.org/fx"
	"jwtAuth/internal/service"
	"jwtAuth/internal/storage"
	"log"
)

var Module = fx.Module("UserService",
	fx.Provide(fx.Annotate(NewDefaultUserService, fx.As(new(service.UserService)))),
)

type DefaultUserService struct {
	storage storage.UserStorage
}

func NewDefaultUserService(lc fx.Lifecycle, storage storage.UserStorage) *DefaultUserService {
	return &DefaultUserService{storage: storage}
}

func (u *DefaultUserService) GetUser(guid string) string {
	log.Println(guid)
	return "token"
}
