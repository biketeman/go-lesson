package analytics

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type Settings struct{
	URL string `json:url`
}
type Analytics struct{
	Settings * Settings `json: settings`
}
type Event struct{
	Username string `json:username`
	Email string `json:email`
}

func New(inputs *Settings)(*Analytics, error){

	var settings *Settings

	settings.URL = inputs.URL

	if inputs.URL == " "{
		inputs.URL = "http://localhost:1337"
	}

	return &Analytics{
		Settings: settings,
	}, nil
}

func (Analytics *Analytics) NewEvent(){

	var event Event
	buffered, _ := json.Marshal(&event)

	_, _ = http.Post("/event", "application", bytes.NewBuffer(buffered) )

}
func (Analytics *Analytics) NewUser(){
	
}