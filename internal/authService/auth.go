package authService

import (
	"errors"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type AuthService struct {
	Secret []byte
	repo   *repository.Repository
}

type Config struct {
	Secret []byte `yaml:"-" env:"QUICK_MENU_SECRET,required"`
}

func New(config Config, repo *repository.Repository) *AuthService {
	return &AuthService{
		Secret: config.Secret,
		repo:   repo,
	}
}

func (as *AuthService) CheckToken(tokenString string) (string, error) {

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return as.Secret, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	userID, exists := claims["user_id"]
	if !exists {
		return "", errors.New("invalid token")
	}

	return userID.(string), nil

}

func (as *AuthService) Login(user model.User) (string, error) {

	userID, err := as.repo.Login(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(as.Secret)

	return tokenString, nil

}
