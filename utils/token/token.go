package token

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenSecret     = "secret"
	TokenExpireTime = 60 * 24 * 90 // 90 days
)

type UserClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) (string, error) {
	expiredAt := time.Now().Add(time.Duration(TokenExpireTime) * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(TokenSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(tokenString string) (*UserClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(TokenSecret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
