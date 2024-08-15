package usecases

import (
	"context"
	"time"

	domain "Task_7/Domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uU *userUsecase) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.CreateUser(ctx, user)
}

func (uU *userUsecase) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.GetUserByUsername(ctx, username)
}

func (uU *userUsecase) GetUserByID(c context.Context, userID string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.GetUserByID(ctx, userID)
}

func (uU *userUsecase) PromoteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.PromoteUser(ctx, userID)
}

func (uU *userUsecase) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.DeleteUser(ctx, userID)
}

func (uU *userUsecase) GetAllUsers(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uU.contextTimeout)
	defer cancel()

	return uU.userRepository.GetAllUsers(ctx)
}
