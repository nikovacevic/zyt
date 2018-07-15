package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nikovacevic/zyt/internal/app/zyt"
)

// Config provides configuration to a JWT service
type Config struct {
	key string
}

// Service is a token service implementing JWTs
type Service struct {
	config *Config
}

// GenerateToken creates and signs a JWT with claims for the given user.
func (s *Service) GenerateToken(user *zyt.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(1440) * time.Hour), // expires in 60 days
	})

	tokenString, err := token.SignedString(s.config.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
