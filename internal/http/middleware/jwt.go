package middleware

import "github.com/golang-jwt/jwt/v4"

type MyCustomClaims struct {
	IdUser int    `json:"iduser"`
	Name   string `json:"name"`
	jwt.StandardClaims
}