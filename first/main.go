package main

import (
	"fmt"

	"./animal"
	"./user"
)

func main() {

	dolphin, _ := animal.NewDolphin(animal.Dolphin{
		Name:   "gang",
		Region: "Australia",
		Age:    12,
	})

	u, _ := user.NewUser(user.User{
		FirstName: "fn",
		LastName:  "ln",
		Age:       12,
	})

	u.AddAnimals([]*animal.Dolphin{
		dolphin,
	})

	fmt.Println(dolphin)
	fmt.Println(u)
}
