package token

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"jwtAuth/internal/storage/postgres"
	"log"
)

var Module = fx.Module("TokenStorage",
	fx.Provide(fx.Annotate(NewDefaultTokenStorage, fx.As(new(Storage)))),
)

type Storage interface {
	GetRefreshToken(guid string) (string, error)
	SaveRefreshToken(guid, refreshToken, sessionId string) (string, error)
}

type DefaultTokenStorage struct {
	storage *postgres.Storage
}

func NewDefaultTokenStorage(lc fx.Lifecycle, storage *postgres.Storage) *DefaultTokenStorage {
	return &DefaultTokenStorage{storage: storage}
}

func (t *DefaultTokenStorage) GetRefreshToken(guid string) (string, error) {
	log.Println("GetRefreshToken not implemented")
	query := "SELECT refresh_token FROM tokens WHERE user_id = $1"
	var refreshToken uuid.UUID
	err := t.storage.Db.QueryRow(query, guid).Scan(&refreshToken)
	if err != nil {
		return "", err
	}
	return refreshToken.String(), nil
}

func (t *DefaultTokenStorage) SaveRefreshToken(guid, refreshToken, sessionId string) (string, error) {
	userUUID, err := uuid.Parse(guid)
	if err != nil {
		return "", err
	}
	sessionUUID, err := uuid.Parse(sessionId)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	bcryptRefresh, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	query := "INSERT INTO tokens(refresh_token, session_id, user_id) VALUES ($1::bytea, $2, $3)"
	_, err = t.storage.Db.Query(query, bcryptRefresh, sessionUUID, userUUID)
	if err != nil {
		return "", err
	}
	return "", nil
}
