package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"not null;type:varchar(255)" json:"last_name"`
	Email     string `gorm:"not null; unique;type:varchar(255)" json:"email"`
	Password  string `gorm:"not null;type:varchar(255)" json:"password"`
	Age       int64  `gorm:"type:int"`
	Username  string `gorm:"not null;unique;type:varchar(255)" json:"username"`
}

type SignUpInput struct {
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	Username             string `json:"username" validate:"required"`
	Age                  int64  `json:"age"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

