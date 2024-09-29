package token

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"jwtAuth/internal/service/user"
	"jwtAuth/internal/storage/token"
	"log"
	"os"
	"strconv"
	"time"
)

var Module = fx.Module("Service",
	fx.Provide(NewJwtTTL),
	fx.Provide(fx.Annotate(NewDefaultTokenService, fx.As(new(Service)))),
)

type Service interface {
	CreateTokenPair(guid, ip string) (string, string, error)
	RefreshTokenPair(access, refresh, ip string) (string, string, error)
}

type DefaultTokenService struct {
	storage     token.Storage
	userService user.Service
	jwtTTL      time.Duration
}

type JwtTTL time.Duration

func NewJwtTTL() JwtTTL {
	seconds, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	log.Println(seconds, "is jwt ttl")
	if err != nil {
		panic("JWT_TTL must be an integer (seconds)")
	}
	return JwtTTL(time.Duration(seconds) * time.Second)
}

func NewDefaultTokenService(storage token.Storage, userService user.Service, ttl JwtTTL) *DefaultTokenService {
	return &DefaultTokenService{storage: storage, userService: userService, jwtTTL: time.Duration(ttl)}
}

func (d *DefaultTokenService) CreateTokenPair(guid, ip string) (string, string, error) {
	sessionId := uuid.New().String()
	access, err := d.createAccessToken(guid, ip, sessionId)
	if err != nil {
		return "", "", err
	}
	refresh := d.createRefreshToken(guid, ip)
	err = d.storage.SaveRefreshToken(guid, refresh, sessionId)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (d *DefaultTokenService) RefreshTokenPair(access, refresh, ip string) (string, string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(access, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	sessionId := claims["session"].(string)
	guid := claims["sub"].(string)
	refreshInDb, err := d.storage.GetRefreshToken(sessionId)
	if err != nil {
		return "", "", err
	}

	if claims["ip"].(string) != ip {
		log.Println("sending warning email")
	}
	log.Println(refresh, string(refreshInDb))
	err = bcrypt.CompareHashAndPassword(refreshInDb, []byte(refresh))
	if err != nil {
		return "", "", errors.New("refresh token invalid")
	}

	newSessionId := uuid.New().String()
	newAccess, err := d.createAccessToken(guid, ip, newSessionId)
	if err != nil {
		return "", "", err
	}
	newRefresh := d.createRefreshToken(guid, ip)

	err = d.storage.UpdateRefreshToken(newRefresh, sessionId, newSessionId)
	if err != nil {
		return "", "", err
	}
	return newAccess, newRefresh, nil
}

func (d *DefaultTokenService) createAccessToken(guid, ip, sessionId string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":     guid,
		"ip":      ip,
		"session": sessionId,
		"exp":     time.Now().Add(d.jwtTTL).Unix(),
		"iat":     time.Now().Unix(),
	})
	signedString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedString, err
}

func (d *DefaultTokenService) createRefreshToken(guid, ip string) string {
	tokenString := ip + " " + time.Now().Format("2006-01-02 15:04:05")
	log.Println(tokenString)
	return base64.StdEncoding.EncodeToString([]byte(tokenString))
}

func (d *DefaultTokenService) parseRefreshToken(refresh string) (string, error) {
	var buffer bytes.Buffer
	_, err := base64.StdEncoding.Decode([]byte(refresh), buffer.Bytes())
	if err != nil {
		return "", err
	}
	return buffer.String(), err
}
