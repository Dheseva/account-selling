package service

import "account-selling/config"

type ItemService struct {
	// Implementasi sesuai kebutuhan Anda
	DB Database
}

func NewItemService(db Database) *ItemService {
    return &ItemService{DB: db}
}

func (db *ItemService) Create(data interface{}) error {
	return config.DB.Create(data).Error
}

func (db *ItemService) Where(query interface{}, args ...interface{}) Database {
	return db // Implementasi sesuai kebutuhan Anda
}

func (db *ItemService) First(out interface{}, where ...interface{}) Database {
	return db // Implementasi sesuai kebutuhan Anda
}