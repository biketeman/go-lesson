package handler

import (
	"database/sql"
	"fmt"

	// ça veut dire que l'on ne va pas s'en servir si on met un underscore devant.
	// Mais par contre on va passer dans la fonction d'init mais on ne va pas se
	// servir de des fonction qu'elle exporte.
	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	var connectionString string

	// on vérifie si on a la variabel d'envirronement.
	// connectionString = os.Getenv("DATABASE_URL")
	// if connectionString == "" {
	// 	return nil, fmt.Errorf("DATABASE_URL not set !")
	// } else {
		// a enlever quand on pousse en prod
		connectionString = "postgres://chrdtcltzbedtq:6ff0bb4d8b971dbec9c729d5816985675b04e861459599b9519534c68ac779b0@ec2-54-247-85-251.eu-west-1.compute.amazonaws.com:5432/d2oh836th98cvr"
	// }

	// Ouvrir la connexion SQL 
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("DB : ", err)
	}

	return db, nil
}
