package user

import (
	"fmt"

	"../animal"
)

type User struct {
	FirstName  string
	LastName   string
	Age        int32
	Email      string
	Dolphins   []*animal.Dolphin
	Properties map[string]interface{}
}

func NewUser(u User) (*User, error) {
	if u.FirstName == "" {
		return &User{}, fmt.Errorf("First name not defined")
	}

	if u.LastName == "" {
		return &User{}, fmt.Errorf("Last name not defined")
	}

	if u.Age < 18 {
		return &User{}, fmt.Errorf("Not young enough")
	}

	return &User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Dolphins:  u.Dolphins,
	}, nil
}

func (user *User) UpdatedFirstName(firstName string) {
	user.FirstName = firstName
}

func (user *User) AddAnimals(dolphins []*animal.Dolphin) {
	user.Dolphins = append(user.Dolphins, dolphins...)
}
