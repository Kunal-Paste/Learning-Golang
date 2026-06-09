package user

import (
	"context"
	"errors"
	"fmt"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResult{}, err
	}

	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("hashing password failed : %w", err)
	}

	now := time.Now().UTC()

	u := User{
		Email:        email,
		PasswordHash: string(hashByte),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	created, err := s.repo.Create(ctx, u)
	if err != nil {
		return AuthResult{}, err
	}

	token, err := auth.CreateToken(s.JWTsecret, created.ID.Hex(), created.Role)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublic(created),
	}, nil

}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.ToLower(strings.TrimSpace(input.Password))

	if email == "" || password == "" {
		return AuthResult{}, errors.New("email & password is required for login")
	}

	if len(password) < 6 {
		return AuthResult{}, errors.New("password must be atleast 6 character long.")
	}

	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return AuthResult{}, errors.New("invalid credentials")
		}
		return AuthResult{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return AuthResult{}, errors.New("invalid credentials or wrong password!")
	}

	token, err := auth.CreateToken(s.JWTsecret, u.ID.Hex(), u.Role)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublic(u),
	}, nil
}
