package repositories

import (
	"errors"

	"github.com/vkhoa145/go-chat-next-ts/app/models"
)

func (r UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	result := r.db.Table("users").Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
