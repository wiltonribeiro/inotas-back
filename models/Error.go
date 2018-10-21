package models

import "fmt"

type Error struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func ErrorResponse(err error, code int) (Error){

	response := Error{
		Code:code,
		Message:fmt.Sprint(err),
	}
	fmt.Println(response)
	return response
}
