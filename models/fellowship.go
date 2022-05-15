package models

import (
	"time"
)


type FellowshipMember struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

type MordorMember struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

//get columns
func (model *FellowshipMember) GetColumns() string {
	return "name, surname, created_at, is_active"
}

//get values
func (model *FellowshipMember) GetValues() string {
	return model.Name + ", " + model.Surname
}
