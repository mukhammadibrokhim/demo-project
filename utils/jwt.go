package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("MySecretMySecretMySecretMySecretMySecret")

type Claims struct {
	Username string `json:"username"`
	Roles    string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(username, roles string) (string, error) {
	claims := &Claims{
		Username: username,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
			Issuer:    username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
