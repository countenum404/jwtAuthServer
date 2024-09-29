package token

import (
	"errors"
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
	GetRefreshToken(sessionId string) ([]byte, error)
	UpdateRefreshToken(newRefresh, oldSessionId, newSessionId string) error
	SaveRefreshToken(guid, refreshToken, sessionId string) error
}

type DefaultTokenStorage struct {
	storage *postgres.Storage
}

func NewDefaultTokenStorage(lc fx.Lifecycle, storage *postgres.Storage) *DefaultTokenStorage {
	return &DefaultTokenStorage{storage: storage}
}

func (t *DefaultTokenStorage) GetRefreshToken(sessionId string) ([]byte, error) {
	log.Println("GetRefreshToken", sessionId)
	query := "SELECT refresh_token FROM tokens WHERE session_id=$1"
	var refreshToken []byte
	err := t.storage.Db.QueryRow(query, sessionId).Scan(&refreshToken)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user hasn't such refresh token")
	}

	return refreshToken, err
}

func (t *DefaultTokenStorage) UpdateRefreshToken(newRefresh, oldSessionId, newSessionId string) error {
	log.Println("UpdateRefreshToken", "newRefresh:", newRefresh, "oldSessionId:", oldSessionId, "newSessionId:", newSessionId)

	queryUpdateRefreshToken := "UPDATE tokens SET refresh_token=$1 WHERE session_id=$2"
	queryUpdateSessionId := "UPDATE tokens SET session_id=$1 WHERE session_id=$2"

	osid, err := uuid.Parse(oldSessionId)
	if err != nil {
		return err
	}
	nsid, err := uuid.Parse(newSessionId)
	if err != nil {
		return err
	}

	bcryptRefresh, err := bcrypt.GenerateFromPassword([]byte(newRefresh), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = t.storage.Db.Exec(queryUpdateRefreshToken, bcryptRefresh, oldSessionId)
	if err != nil {
		return err
	}

	_, err = t.storage.Db.Exec(queryUpdateSessionId, nsid, osid)
	if err != nil {
		return err
	}

	return err
}

func (t *DefaultTokenStorage) SaveRefreshToken(guid, refreshToken, sessionId string) error {
	log.Println("refreshToken", refreshToken)
	query := "INSERT INTO tokens(refresh_token, session_id, user_id) VALUES ($1, $2, $3)"

	userUUID, err := uuid.Parse(guid)
	if err != nil {
		return err
	}

	sessionUUID, err := uuid.Parse(sessionId)
	if err != nil {
		return err
	}

	bcryptRefresh, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = t.storage.Db.Query(query, bcryptRefresh, sessionUUID, userUUID)
	if err != nil {
		return err
	}

	return err
}
