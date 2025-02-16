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
	signingKey = "fjlskjJISJofmdslkijou43298742"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user types.SignInInput) (int, error) {
	const op = "AuthService.CreateUser"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call CreateUser")

	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(user types.UserDAO) (string, error) {
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
	if accessToken == "" {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(accessToken, &types.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("malformed token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token has expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token is not valid yet")
			} else {
				return nil, errors.New("token validation failed")
			}
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(*types.TokenClaims)
	if !ok {
		return nil, errors.New("failed to cast token claims")
	}

	return claims, nil
}

func (s *AuthService) CheckAuthData(username, password string) (types.UserDAO, error) {
	return s.repo.GetUser(username, s.generatePasswordHash(password))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
