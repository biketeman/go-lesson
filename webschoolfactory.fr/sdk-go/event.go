package analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (analytics *Analytics) NewEvent(event *Event) error {

	if event.Name == "" {
		return fmt.Errorf("Event name not set")
	}

	if event.UserID == "" {
		return fmt.Errorf("User ID not set")
	}

	buffered, _ := json.Marshal(event)

	url := analytics.Settings.URL + "/event"
	_, _ = http.Post(url, "application/json", bytes.NewBuffer(buffered))

	return nil
}
