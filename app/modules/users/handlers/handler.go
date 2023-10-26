package handlers

import (
	userRepo "github.com/vkhoa145/go-chat-next-ts/app/modules/users/repositories"
	userUsecase "github.com/vkhoa145/go-chat-next-ts/app/modules/users/usecase"
)

type UserHandler struct {
	userRepo    userRepo.UserRepoInterface
	userUsecase userUsecase.UserUsecaseInterface
}

func NewUserHandler(userRepo userRepo.UserRepoInterface, userUsecase userUsecase.UserUsecaseInterface) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		userUsecase: userUsecase,
	}
}
