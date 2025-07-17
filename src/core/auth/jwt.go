package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateTokens(userID int32, roleID int32) (access, refresh string, err error) {
	access, err = generate(userID, roleID, 15*time.Minute)
	if err != nil {
		return
	}
	refresh, err = generate(userID, roleID, 7*24*time.Hour)
	return
}

func generate(uid, rid int32, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  uid,
		"role": rid,
		"exp":  time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
