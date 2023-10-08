package repository

import "account-selling/internal/entity"

type UserDataRepository interface {
	CreateUserData(userData *entity.UserData) error
	FindById(id uint) (*entity.UserData, error)
}