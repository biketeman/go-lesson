package main

import analytics "webschoolfactory.fr/sdk-go"

func main() {
	client, _ := analytics.New(&analytics.Settings{
		// URL: "http:"
	})

	// client.NewUser(&analytics.User{
	// 	Username: "arthur",
	// })

	client.NewEvent(&analytics.Event{
		Name:   "Super event ter",
		UserID: "1IwsZTolFQM12Ynen248xrM8yFL",
	})
}
