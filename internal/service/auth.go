package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/types"
)

const (
	salt       = "jldsajlf%sfldj#dfsf"
	signingKey = "ajfas#user#ajldfj32"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user types.UserType) (int, error) {
	const op = "AuthService.CreateUser"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call CreateUser")

	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(user types.UserType) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParserToken(accessToken string) (*types.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &types.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		if _, ok := token.Claims.(*types.TokenClaims); ok {

			return []byte(signingKey), nil
		}

		return nil, errors.New("token claims are not of type *tokenClaims")
	})

	if err != nil {
		slog.Info("Invalid token", err)
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("токен невалиден")
	}

	claims, _ := token.Claims.(*types.TokenClaims)

	return claims, nil
}

func (s *AuthService) CheckAuthData(username, password string) (types.UserType, error) {
	return s.repo.GetUser(username, s.generatePasswordHash(password))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
