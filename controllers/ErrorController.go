package controllers

import "log"

func CheckFail(err error){
	if err != nil {
		log.Fatal(err)
	}
}