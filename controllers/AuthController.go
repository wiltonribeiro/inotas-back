package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"inotas-back/enviroment"
)

type AuthController struct {
	Email string
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (controller AuthController) GenerateAuth() (tokenString string, err error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": controller.Email,
	})

	mySigningKey := enviroment.SecretKey
	tokenString, err = token.SignedString(mySigningKey)

	return
}

func (controller AuthController) CheckAuth(tokenString string) (err error){

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return enviroment.SecretKey, nil
	})

	if _, ok := token.Claims.(*Claim); !(ok && token.Valid) {
		return err
	}
	return nil
}
