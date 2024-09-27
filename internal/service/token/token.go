package token

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"jwtAuth/internal/service"
	"jwtAuth/internal/storage"
	"log"
	"os"
	"time"
)

var Module = fx.Module("TokenService",
	fx.Provide(fx.Annotate(NewDefaultTokenService, fx.As(new(service.TokenService)))),
)

type DefaultTokenService struct {
	storage     storage.TokenStorage
	userService service.UserService
}

func NewDefaultTokenService(lc fx.Lifecycle, storage storage.TokenStorage, userService service.UserService) *DefaultTokenService {
	return &DefaultTokenService{storage: storage, userService: userService}
}

func (d *DefaultTokenService) CreateTokenPair(guid, ip string) (string, string, error) {
	sessionId := uuid.New().String()
	access, err := d.createAccessToken(guid, ip, sessionId)
	if err != nil {
		return "", "", err
	}
	refresh := d.createRefreshToken(guid, ip)
	_, err = d.storage.SaveRefreshToken(guid, refresh, sessionId)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (d *DefaultTokenService) GetRefreshToken(guid string) (string, error) {
	log.Println("GetRefreshToken not implemented yet")
	return "", nil
}

func (d *DefaultTokenService) createAccessToken(guid, ip, sessionId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":     guid,
		"ip":      ip,
		"session": sessionId,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
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
