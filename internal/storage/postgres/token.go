package postgres

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type TokenStorage struct {
	storage *Storage
}

func NewTokenStorage(lc fx.Lifecycle, storage *Storage) *TokenStorage {
	return &TokenStorage{storage: storage}
}

func (t *TokenStorage) GetRefreshToken(guid string) (string, error) {
	log.Println("GetRefreshToken not implemented")
	query := "SELECT refresh_token FROM tokens WHERE user_id = $1"
	var refreshToken uuid.UUID
	err := t.storage.db.QueryRow(query, guid).Scan(&refreshToken)
	if err != nil {
		return "", err
	}
	return refreshToken.String(), nil
}

func (t *TokenStorage) SaveRefreshToken(guid, refreshToken, sessionId string) (string, error) {
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
	_, err = t.storage.db.Query(query, bcryptRefresh, sessionUUID, userUUID)
	if err != nil {
		return "", err
	}
	return "", nil
}
