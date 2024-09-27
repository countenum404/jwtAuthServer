package user

import (
	"go.uber.org/fx"
	"jwtAuth/internal/storage/user"
	"log"
)

var Module = fx.Module("UserService",
	fx.Provide(fx.Annotate(NewDefaultUserService, fx.As(new(Service)))),
)

type Service interface {
	GetUser(guid string) string
}

type DefaultUserService struct {
	storage user.Storage
}

func NewDefaultUserService(lc fx.Lifecycle, storage user.Storage) *DefaultUserService {
	return &DefaultUserService{storage: storage}
}

func (u *DefaultUserService) GetUser(guid string) string {
	log.Println(guid)
	return "token"
}
