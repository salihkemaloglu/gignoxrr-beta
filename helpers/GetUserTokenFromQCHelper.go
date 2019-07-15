package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
)

// Item ...
type Item struct {
	Result string `bson:"result" json:"result"`
	Status string `bson:"status" json:"status"`
}

//GetUserToken ...
func GetUserToken(user repo.User) (string, error) {
	url := "http://localhost:8902/login"

	userJSON, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return item.Result, err
}
