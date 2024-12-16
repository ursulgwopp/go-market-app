package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/market-api/internal/models"
	"github.com/ursulgwopp/market-api/internal/repository"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(req models.SignUpRequest) (int, error) {
	if err := validateUsername(req.Username); err != nil {
		return -1, err
	}

	if err := validatePassword(req.Password); err != nil {
		return -1, err
	}

	if err := validateEmail(req.Email); err != nil {
		return -1, err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.SignUp(req)
}

func (s *AuthService) GenerateToken(req models.SignInRequest) (string, error) {
	exists, err := s.repo.CheckUsernameExists(req.Username)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("username does not exists")
	}

	req.Password = generatePasswordHash(req.Password)
	UserId, err := s.repo.SignIn(req)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
