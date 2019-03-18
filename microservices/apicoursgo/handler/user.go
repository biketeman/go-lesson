package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/segmentio/ksuid"
)

type User struct {
	ID        ksuid.KSUID
	Username  string
	Email     string
	CreatedAt time.Time
}

type PayloadUser struct {
	// tag sur les structure. on a la key (username), type (string) et tags (`json:"username"`).
	// Le tag va etre utilisé par le paquet JSON.
	Username string `json:"username"`
	Email    string `json:"email"`
}

func addUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)

	var user PayloadUser
	// rentré dans le pointeur de type user ce que json a décodé.
	err := decoder.Decode(&user)
	if err != nil {
		res.WriteHeader(400)
		return
	}

	fmt.Println("USERNAME", user.Username)
	fmt.Println("EMAIL", user.Email)

	db, err := NewDatabase()
	// fermer la connection quand la fonction s'arrete
	defer db.Close()

	// verification de l'erreur si jamais NewDatabase nous en renvoie une.
	if err != nil {
		res.WriteHeader(500)
		return
	}

	// Création d'un id unique
	id := ksuid.New()

	// Prend des variadic functions en parametre : Fonction avec un nombre indéterminé d'arguments
	_, err = db.Query("INSERT INTO users (id, username, email) VALUES ($1, $2, $3);",
		id.String(), user.Username, user.Email)

	// Verfiication de l'erreur lors de la queries
	if err != nil {
		res.WriteHeader(400)
		return
	}

	body := []byte(`{"Route": "Post User"}`)
	// Permet d'écrire le status code
	res.WriteHeader(201)
	// ecrire des Bytes en tant que réponse
	res.Write(body)
}
