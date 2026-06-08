package user

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	repo      *Repo
	JWTsecret string
}

func NewService(repo *Repo, JWTsecret string) *Service {
	return &Service{
		repo:      repo,
		JWTsecret: JWTsecret,
	}
}

type RegisterInput struct {
	Email    string `josn:"email"`
	Password string `json:"paassword"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.ToLower(strings.TrimSpace(input.Password))

	if email == "" || password == "" {
		return AuthResult{}, errors.New("email and password is required")
	}

	if len(password) < 6 {
		return AuthResult{}, errors.New("password must be atleast 6 character long")
	}

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("email is already registered, so please try with different email")
	}

	if err != nil || errors.Is(err, mongo.ErrNilDocument) {
		return AuthResult{}, err
	}
}
