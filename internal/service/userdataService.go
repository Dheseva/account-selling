package service

import (
	"account-selling/internal/entity"
	"account-selling/internal/repository"
	"time"
)

type RegisterUserDataUsecase struct {
	UserDataRepository repository.UserDataRepository
}

func NewRegisterUserdataUseCase(userdata repository.UserDataRepository) *RegisterUserDataUsecase {
	return &RegisterUserDataUsecase{
		UserDataRepository: userdata,
	}
}

func (uc *RegisterUserDataUsecase) Execute(userdata *entity.UserData, data map[string]string) (*entity.UserData,error) {
	userdata.Nickname = data["name"]
	userdata.Created_at = time.Now().UnixMilli()
	userdata.Updated_at = time.Now().UnixMilli()
	if err := uc.UserDataRepository.CreateUserData(userdata); err != nil {
		return nil , err
	}
	return userdata, nil
}