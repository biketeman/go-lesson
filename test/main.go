package main

import (
	"fmt"

	"./animals"
)

func main() {
	benjamin, _ := animals.Addanimals(animals.Animal{
		Name: "Benjamin",
		Age:  14,
	})
	giles, _ := animals.Addanimals(animals.Animal{
		Name: "Gilles",
		Age:  15,
	})
	fmt.Println(benjamin)
	fmt.Println(giles)

}
