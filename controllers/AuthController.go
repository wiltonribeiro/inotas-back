package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"inotas-back/enviroment"
	"fmt"
	"inotas-back/models"
)

type AuthController struct {

}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (controller AuthController) GenerateAuth(email string) (tokenString string, error models.Error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	mySigningKey := enviroment.SecretKey
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		error = models.ErrorResponse(err, 500)
	}

	return
}

func (controller AuthController) CheckAuth(tokenString string) (string, models.Error){

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return enviroment.SecretKey, nil
	})

	if err != nil {
		return "", models.ErrorResponse(err, 401)
	}

	claims, ok := token.Claims.(*Claim)
	if ok && token.Valid {
		return fmt.Sprint(claims.Email), models.Error{}
	} else {
		return "", models.ErrorResponse(err, 403)
	}
}
