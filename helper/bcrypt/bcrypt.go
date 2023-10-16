package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(pass string) ([]byte, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	return password, err
}

func ValidateHash(up []byte,pass []byte) error {
	if err := bcrypt.CompareHashAndPassword(up, pass); err != nil{
		return err
	}
	return nil
}