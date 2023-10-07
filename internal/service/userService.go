package service

import (
	"account-selling/config"
	"account-selling/internal/entity"
	"account-selling/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	UserRepository repository.UserRepository
	UserDataRepository repository.UserDataRepository
}

func NewRegisterUserUseCase(userRepo repository.UserRepository, userdataRepo repository.UserDataRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepo,
		UserDataRepository: userdataRepo,
	}
}

func (uc *RegisterUserUseCase) Execute(user *entity.User, userdata *entity.UserData, data map[string]string) error {

	userdata.Nickname = data["name"]
	userdata.Created_at = time.Now().UnixMilli()
	userdata.Updated_at = time.Now().UnixMilli()
	if err := config.DB.Create(&userdata).Error; err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user.Name = data["name"]
	user.Password = password
	user.UData_id = int(userdata.Id)
	user.Lastlogin = time.Now().UnixMilli()
	user.Created_at = time.Now().UnixMilli()
	user.Updated_at = time.Now().UnixMilli()
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return uc.UserRepository.Create(user)
}