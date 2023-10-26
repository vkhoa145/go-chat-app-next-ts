package usecase

import (
	"errors"

	"github.com/vkhoa145/go-chat-next-ts/app/models"
)

func (u UserUsecase) SignUp(payload *models.SignUpInput) (*models.User, error) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, errors.New("Password does not match")
	}

	createUser, err := u.userRepo.CreateUser(payload)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}
