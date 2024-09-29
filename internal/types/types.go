package types

import "github.com/google/uuid"

type User struct {
	Id       uuid.Domain
	Name     string
	Username string
	Email    string
	Password string
}

type RefreshRequest struct {
	Access  string `json:"Access"`
	Refresh string `json:"Refresh"`
}

type RefreshContext struct {
	Guid            string
	OldRefresh      string
	NewRefreshToken string
	OldSession      string
	NewSession      string
}
