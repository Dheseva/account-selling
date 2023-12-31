package repository

import "account-selling/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
