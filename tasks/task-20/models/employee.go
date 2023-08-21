package models

import (
	//"encoding/json"
	"encoding/json"
	"errors"
)

type Employee struct {
	Id string
	Name string `json:"name`
	Email string `json:email`
	Mobile string `json:mobile`
	Status string `json:status`
	LastModified int64 `json:lastModified`
}

func (e *Employee)Validate() error{
	if e.Name=="" || e.Id=="" || e.Email=="" || e.Mobile=="" || e.Status==""{
		return errors.New("Invalid input type]")
	}
	return nil
}

func (e Employee)ToBytes()([]byte,error){
	return json.Marshal(e)
}
