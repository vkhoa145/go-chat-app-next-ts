package usecase

import (
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	userRepo "github.com/vkhoa145/go-chat-next-ts/app/modules/users/repositories"
)

type UserUsecaseInterface interface {
	SignUp(payload *models.SignUpInput) (*models.User, error)
	SignIn(payload *models.SignInInput) (*models.User, error)
}

type UserUsecase struct {
	userRepo userRepo.UserRepoInterface
}

func NewUserUsecase(userRepo userRepo.UserRepoInterface) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
