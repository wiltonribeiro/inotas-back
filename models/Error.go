package models

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
)

type Error struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Time time.Time `json:"time"`
}

func ErrorResponse(err error, code int) (Error){
	response := Error{Code:code, Message:fmt.Sprint(err), Time: time.Now()}
	response.writeLog()
	return response
}

func (error Error) writeLog(){
	exits := error.existFile()
	f, err := os.OpenFile("errors.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		data,_ := json.Marshal(error)
		if exits {
			f.Write([]byte(","))
			f.Write([]byte(data))
		} else {
			f.Write([]byte(data))
		}
		defer  f.Close()
	}
}

func (error Error) existFile() bool{
	if _, err := os.Stat("errors.json"); os.IsNotExist(err) {
		return false
	}
	return true
}
