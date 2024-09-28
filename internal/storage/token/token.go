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
	GetRefreshToken(guid, sessionId string) (string, []byte, error)
	GetSessionIdByRefresh(refresh string) (string, string, error)
	UpdateRefreshToken(guid, refresh, sessionId, newSessionId string) error
	SaveRefreshToken(guid, refreshToken, sessionId string) error
}

type DefaultTokenStorage struct {
	storage *postgres.Storage
}

func NewDefaultTokenStorage(lc fx.Lifecycle, storage *postgres.Storage) *DefaultTokenStorage {
	return &DefaultTokenStorage{storage: storage}
}

func (t *DefaultTokenStorage) GetRefreshToken(guid, sessionId string) (string, []byte, error) {
	log.Println("GetRefreshToken", guid, sessionId)
	query := "SELECT id, refresh_token FROM tokens WHERE user_id=$1 AND session_id=$2"

	var id uuid.UUID
	var refreshToken []byte
	err := t.storage.Db.QueryRow(query, guid, sessionId).Scan(&id, &refreshToken)
	if err != nil {
		return "", nil, errors.New("user hasn't such refresh token")
	}

	return id.String(), refreshToken, err
}

func (t *DefaultTokenStorage) GetSessionIdByRefresh(refresh string) (string, string, error) {
	log.Println("GetSessionIdByRefresh", refresh)

	_, err := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	query := "SELECT user_id, session_id FROM tokens WHERE refresh_token=decode($1, 'hex')"

	var sessionId, guid uuid.UUID
	err = t.storage.Db.QueryRow(query, refresh).Scan(&guid, &sessionId)
	if err != nil {
		return "", "", err
	}

	return sessionId.String(), guid.String(), err
}

func (t *DefaultTokenStorage) UpdateRefreshToken(guid, refresh, sessionId, newSessionId string) error {
	log.Println("UpdateRefreshToken", guid, refresh, sessionId, newSessionId)

	query := "UPDATE tokens SET refresh_token=$1, session_id=$2 WHERE user_id=$3 AND session_id=$4"

	ssid, err := uuid.Parse(sessionId)
	if err != nil {
		return err
	}

	userId, err := uuid.Parse(guid)
	if err != nil {
		return err
	}

	bcryptRefresh, err := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newSsid, err := uuid.Parse(sessionId)
	if err != nil {
		return err
	}

	_, err = t.storage.Db.Exec(query, bcryptRefresh, newSsid, userId, ssid)
	if err != nil {
		return err
	}
	return nil
}

func (t *DefaultTokenStorage) SaveRefreshToken(guid, refreshToken, sessionId string) error {
	query := "INSERT INTO tokens(refresh_token, session_id, user_id) VALUES ($1::bytea, $2, $3)"

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
