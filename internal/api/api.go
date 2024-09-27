package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	fx.Provide(
		NewAddr,
		fx.Annotate(NewAuthHandlers, fx.As(new(Authenticator))),
	),
	fx.Invoke(NewJwtAuthApi),
)

type Addr string

func NewAddr() Addr {
	return ":8080"
}

type JwtAuthApi struct {
	addr     Addr
	handlers Authenticator
}

func (j *JwtAuthApi) Run() error {
	router := gin.Default()
	router.GET("/ping", pong)

	router.POST("/auth", j.handlers.HandleCreateJWT)
	router.PUT("/auth", j.handlers.HandleRefreshJWT)

	err := router.Run(string(j.addr))
	if err != nil {
		return err
	}
	return nil
}

func NewJwtAuthApi(lc fx.Lifecycle, addr Addr, handlers Authenticator) *JwtAuthApi {
	j := &JwtAuthApi{addr: addr, handlers: handlers}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := func() error {
					err := j.Run()
					if err != nil {

						return err
					}
					return nil
				}()
				if err != nil {
				}
			}()
			return nil
		},
	})
	return j
}
