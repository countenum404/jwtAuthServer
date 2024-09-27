package service

type UserService interface {
	GetUser(guid string) string
}

type TokenService interface {
	CreateTokenPair(guid, ip string) (string, string, error)
	GetRefreshToken(guid string) (string, error)
}
