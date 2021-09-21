package models

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID   int32
	Username string
	*jwt.StandardClaims
}
