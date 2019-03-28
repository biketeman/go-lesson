package analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (analytics *Analytics) NewUser(user *User) error {

	if user.Username == "" {
		return fmt.Errorf("Username not set")
	}

	if user.Email == "" {
		return fmt.Errorf("Email not set")
	}

	buffered, _ := json.Marshal(user)

	url := analytics.Settings.URL + "/user"
	_, _ = http.Post(url, "application/json", bytes.NewBuffer(buffered))

	return nil
}
