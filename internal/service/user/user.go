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
	GetUserEmailById(guid string) (string, error)
}

type DefaultUserService struct {
	storage user.Storage
}

func (u *DefaultUserService) GetUserEmailById(guid string) (string, error) {
	id, err := u.storage.GetUserEmailById(guid)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewDefaultUserService(lc fx.Lifecycle, storage user.Storage) *DefaultUserService {
	return &DefaultUserService{storage: storage}
}

func (u *DefaultUserService) GetUser(guid string) string {
	log.Println(guid)
	return "token"
}
