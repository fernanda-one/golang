package service

import (
	"github.com/fernanda-one/golang_api/Dto"
	"github.com/fernanda-one/golang_api/entities"
	"github.com/fernanda-one/golang_api/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	Update(user Dto.UserUpdateDto) entities.User
	Profile(userID string) entities.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user Dto.UserUpdateDto) entities.User {
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updateduser := service.userRepository.UpdateUser(userToUpdate)
	return updateduser
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}
