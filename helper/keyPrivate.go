package helper

import "os"

func PrivateKey() string {
	var Secretkey = os.Getenv("PRIVATE_KEY_JWT")
	return Secretkey
}