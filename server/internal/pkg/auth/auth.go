package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var tokenLifetime = time.Duration(100) * time.Hour

type AuthManager struct {
	signingKey []byte
}

func NewAuthManager(signingKey []byte) *AuthManager {
	return &AuthManager{signingKey: signingKey}
}

func (a *AuthManager) MakeAuth(UserId uint) (string, error) {
	expTime := time.Now().Add(tokenLifetime)
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(UserId),
		ExpiresAt: jwt.NewNumericDate(expTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (a *AuthManager) FetchAuth(tokenString string) (*map[string]string, error) {
	claims := jwt.RegisteredClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return a.signingKey, nil
	})
	if err != nil {
		return &map[string]string{}, err
	}
	if !tkn.Valid {
		return &map[string]string{}, errors.New("Invalid Token")
	}
	return ConvertClaimsToMap(&claims), nil
}

func ConvertClaimsToMap(claims *jwt.RegisteredClaims) *map[string]string {
	sub := claims.Subject
	exp := claims.ExpiresAt
	res := map[string]string{
		"sub": sub,
		"exp": fmt.Sprint(exp),
	}
	return &res
}
