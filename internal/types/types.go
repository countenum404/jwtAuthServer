package types

import "github.com/google/uuid"

type User struct {
	Id       uuid.Domain
	Name     string
	Username string
	Email    string
	Password string
}
