package security

import (
	"strings"
	"time"

	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/PauloPHAL/microservices/pkg/perrors"
	"github.com/golang-jwt/jwt/v4"
)

type TokenGenerator struct {
	jwtSecret []byte
}

func NewTokenGenerator(secret string) interfaces.TokenGenerator {
	return &TokenGenerator{
		jwtSecret: []byte(secret),
	}
}

func (t *TokenGenerator) GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"iat":  time.Now().Unix(),
		"type": "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.jwtSecret)
	if err != nil {
		return "", perrors.ErrTokenGeneration
	}

	return tokenString, nil
}

func (t *TokenGenerator) GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(t.jwtSecret)
	if err != nil {
		return "", perrors.ErrTokenGeneration
	}

	return tokenString, nil
}

func (t *TokenGenerator) ValidateToken(tokenString, expectedType string) (string, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, perrors.ErrUnexpectedSigning
		}
		return t.jwtSecret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			return "", perrors.ErrTokenExpired
		}
		return "", perrors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", perrors.ErrInvalidTokenClaims
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != expectedType {
		return "", perrors.ErrInvalidToken
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", perrors.ErrInvalidTokenSubject
	}

	return sub, nil
}
