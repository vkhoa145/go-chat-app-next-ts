package repositories

import (
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(data *models.SignUpInput) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}
