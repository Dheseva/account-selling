package service

import (
	bcrypt "account-selling/helper/bcrypt"
	"account-selling/internal/entity"
	"account-selling/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	UserRepository repository.UserRepository
}

func NewRegisterUserUseCase(userRepo repository.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepo,
	}
}

func (uc *RegisterUserUseCase) Execute(user *entity.User, userdata *entity.UserData, data map[string]string) error {

	password, _ := bcrypt.PasswordHash(data["password"])
	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user.Name = data["name"]
	user.Password = password
	user.UData_id = int(userdata.Id)
	user.Lastlogin = time.Now().UnixMilli()
	user.Created_at = time.Now().UnixMilli()
	user.Updated_at = time.Now().UnixMilli()
	if err := uc.UserRepository.Create(user); err != nil {
		return err
	}

	return nil
}