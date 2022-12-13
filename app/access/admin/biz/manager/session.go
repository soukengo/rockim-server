package manager

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.RegisteredClaims
	User SessionUser `json:"user"`
}

type SessionUser struct {
	Name      string
	AvatarUrl string
}
