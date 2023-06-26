package model

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Name        string
	Description string
}

type InputObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (i *InputObject) ToObject() *Object {
	return &Object{
		Name:        i.Name,
		Description: i.Description,
		Model:       gorm.Model{},
	}
}
