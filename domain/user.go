package domain

import (
	"github.com/Kamva/mgm/v2"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
}

type UserRepository interface {
	Echo(user User) User
	Create(user User) *User
	GetUserById(id string) *User
	GetUsers() []User
}

type UserUseCase interface {
	Echo(user User) User
	Create(user User) *User
	GetUserById(id string) *User
	GetUsers() []User
}
