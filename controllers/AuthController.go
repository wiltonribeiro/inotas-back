package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"inotas-back/enviroment"
	"fmt"
)

type AuthController struct {

}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (controller AuthController) GenerateAuth(email string) (tokenString string, err error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	mySigningKey := enviroment.SecretKey
	tokenString, err = token.SignedString(mySigningKey)

	return
}

func (controller AuthController) CheckAuth(tokenString string) (email string,err error){

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return enviroment.SecretKey, nil
	})

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return fmt.Sprint(claims.Email), nil
	}

	return
}
