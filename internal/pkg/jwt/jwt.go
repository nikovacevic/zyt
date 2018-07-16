package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nikovacevic/zyt/internal/app/zyt"
	"github.com/spf13/viper"
)

// Config provides configuration to a JWT service
type Config struct {
	Key string `mapstructure:"jwt_key"`
}

// Service is a token service implementing JWTs
type Service struct {
	key string
}

// GetConfig finds a config file in the given path and constructs a Config
func GetConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cnf Config
	if err := v.Unmarshal(&cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}

// NewService ...TODO
func NewService(cnf Config) *Service {
	return &Service{cnf.Key}
}

// GenerateToken creates and signs a JWT with claims for the given user.
func (s *Service) GenerateToken(user *zyt.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(1440) * time.Hour), // expires in 60 days
	})

	tokenString, err := token.SignedString([]byte(s.key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the signature of a given token. Returns nil if verified,
// an error if not.
func (s *Service) VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{
		"iat": nil,
		"exp": nil,
	}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.key), nil
	})

	if !token.Valid {
		return fmt.Errorf("Token is not valid")
	}

	return err
}
