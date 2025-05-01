package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input RegisterInput) (*User, error)
	GetByEmail(email string) (*User, error)
	ComparePassword(hashedPwd string, plainPwd string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// RegisterInput defines input for registering a user
type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

func (s *userService) Register(input RegisterInput) (*User, error) {
	existingUser, _ := s.repo.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) ComparePassword(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
