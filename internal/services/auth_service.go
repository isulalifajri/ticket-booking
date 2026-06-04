package services

import (
	"errors"

	"ticket-booking/internal/dto"
	"ticket-booking/internal/models"
	"ticket-booking/internal/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(
	userRepo *repositories.UserRepository,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(
	req dto.RegisterRequest,
) error {

	_, err := s.userRepo.FindByEmail(req.Email)

	if err == nil {
		return errors.New("email already registered")
	}

	if err != gorm.ErrRecordNotFound {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "customer",
	}

	return s.userRepo.Create(&user)
}

func (s *AuthService) Login(
	req dto.LoginRequest,
) (*models.User, error) {

	user, err := s.userRepo.FindByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}