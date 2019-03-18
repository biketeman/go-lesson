package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/segmentio/ksuid"
)

type Event struct {
	ID        ksuid.KSUID
	User      *User
	Name      string
	librairy  *Library
	CreatedAt time.Time
}

type PayloadEvent struct {
	// tag sur les structure. on a la key (username), type (string) et tags (`json:"username"`).
	// Le tag va etre utilisé par le paquet JSON.
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func addEvent(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)

	var body ResponsePostEvent

	var event PayloadEvent
	// rentré dans le pointeur de type event ce que json a décodé.
	err := decoder.Decode(&event)
	// erreur de décodage
	if err != nil {
		// on créé notre reponse que l'on va renvoyé.
		body = ResponsePostEvent{
			StatusCode: 400,
			Message:    "Bad Data",
			Data:       nil,
		}

		bytes, _ := json.Marshal(&body)
		res.WriteHeader(400)
		res.Write(bytes)

		return
	}

	fmt.Println("Name : ", event.Name)
	fmt.Println("User_id : ", event.UserID)

	db, err := NewDatabase()
	// fermer la connection quand la fonction s'arrete
	defer db.Close()

	// verification de l'erreur si jamais NewDatabase nous en renvoie une.
	if err != nil {
		// on créé notre reponse que l'on va renvoyé.
		body = ResponsePostEvent{
			StatusCode: 500,
			Message:    "Internal Server Error",
			Data:       nil,
		}

		bytes, _ := json.Marshal(&body)
		res.WriteHeader(500)
		res.Write(bytes)

		return
	}

	// Création d'un id unique
	id := ksuid.New()

	// Prend des variadic functions en parametre : Fonction avec un nombre indéterminé d'arguments
	_, err = db.Query("INSERT INTO events (id, name, user_id) VALUES ($1, $2, $3);",
		id.String(), event.Name, event.UserID)

	// Verification de l'erreur pour la DB.
	if err != nil {
		// on créé notre reponse que l'on va renvoyé.
		body = ResponsePostEvent{
			StatusCode: 400,
			Message:    "Bad data",
			Data:       nil,
		}

		bytes, _ := json.Marshal(&body)
		res.WriteHeader(400)
		res.Write(bytes)
		return
	}

	// on créé notre reponse que l'on va renvoyé.
	body = ResponsePostEvent{
		StatusCode: 201,
		Message:    "Created",
		Data: &ResponsePostEventData{
			Name: event.Name,
		},
	}

	bytes, _ := json.Marshal(&body)

	fmt.Println(err)

	// body = []byte(`{"Route": "Post Event"}`)
	// Permet d'écrire le status code
	res.WriteHeader(201)
	// ecrire des Bytes en tant que réponse
	res.Write(bytes)
}
