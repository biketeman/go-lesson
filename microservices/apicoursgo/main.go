package main

import (
	"microservices/apicoursgo/handler"
	"fmt"
)

func init() {
	// On appelle la fonctiond e database pour savoir si elle fonctionne.
	_, err := handler.NewDatabase()
	// Si il y a une erreur, on affiche l'erreur.
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	service, err := handler.NewHandler(handler.Settings{
		Name:    "cours",
		Address: ":1337",
	})

	if err != nil {
		fmt.Errorf("Error initializing service :", err)
	}

	if err := service.Init(); err != nil {
		fmt.Errorf("Error", err)
	}

	if err := service.Run(); err != nil {
		fmt.Errorf("Error", err)
	}
}
