package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"jwtAuth/internal/service"
	"net/http"
)

type Authenticator interface {
	HandleCreateJWT(*gin.Context)
	HandleRefreshJWT(*gin.Context)
}

type AuthHandlers struct {
	service service.UserService
}

func (a *AuthHandlers) HandleCreateJWT(context *gin.Context) {
	//TODO implement me

}

func (a *AuthHandlers) HandleRefreshJWT(context *gin.Context) {
	//TODO implement me

}

func NewAuthHandlers(lc fx.Lifecycle, service service.UserService) *AuthHandlers {
	return &AuthHandlers{service: service}
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
