package main

import (
	"fmt"
	"sync"
	"time"
)

func Verify(wg *sync.WaitGroup, user *User) {
	time.Sleep(2 * time.Second)
	fmt.Println("User", user.Name, "is verified")
	defer wg.Done()
}

func main() {
	var userList = []*User{
		&User{
			Name: "toto",
			Age:  56,
		},
		&User{
			Name: "pablo",
			Age:  45,
		},
	}

	var wg = sync.WaitGroup{}

	for _, u := range userList {
		wg.Add(1)
		go Verify(&wg, u)
	}
	wg.Wait()
	fmt.Println("Done")
}
