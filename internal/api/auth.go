package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"jwtAuth/internal/service/token"
	"jwtAuth/internal/service/user"
	"net/http"
)

type Authenticator interface {
	HandleCreateJWT(*gin.Context)
	HandleRefreshJWT(*gin.Context)
}

type AuthHandlers struct {
	userService  user.Service
	tokenService token.Service
}

func NewAuthHandlers(lc fx.Lifecycle, userService user.Service, tokenService token.Service) *AuthHandlers {
	return &AuthHandlers{userService: userService, tokenService: tokenService}
}

func (a *AuthHandlers) HandleCreateJWT(context *gin.Context) {
	access, refresh, err := a.tokenService.CreateTokenPair(context.Query("guid"), context.ClientIP())
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Access": access, "Refresh": refresh})
}

func (a *AuthHandlers) HandleRefreshJWT(context *gin.Context) {
	// To come up with workflow
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
