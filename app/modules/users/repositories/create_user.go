package repositories

import (
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (r UserRepo) CreateUser(data *models.SignUpInput) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return nil, err
	}

	var newUser = &models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Age:       data.Age,
		Password:  string(hashPassword),
	}

	result := r.db.Table("users").Create(newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}
