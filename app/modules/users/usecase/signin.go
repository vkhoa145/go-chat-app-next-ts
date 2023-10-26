package usecase

import (
	"errors"

	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) SignIn(payload *models.SignInInput) (*models.User, error) {
	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if errPass != nil {
		return nil, errors.New("password is not correct")
	}

	return user, nil
}
