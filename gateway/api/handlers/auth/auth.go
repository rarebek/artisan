package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager struct {
	secretKey string
}

func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{secretKey: secretKey}
}

func (tm *TokenManager) GenerateToken(objectID string) (string, error) {
	claims := jwt.MapClaims{
		"id":  objectID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tm.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (tm *TokenManager) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, exists := claims["id"].(string); exists && id != "" {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).After(time.Now()) {
					return true, nil
				}
			}
		}
	}

	return false, fmt.Errorf("invalid or expired token")
}

func (tm *TokenManager) ExtractIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, exists := claims["id"].(string); exists && id != "" {
			return id, nil
		}
	}

	return "", fmt.Errorf("invalid token")
}
