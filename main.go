package clearing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers() ([]User, error) {
	url := "http://localhost:8080/users"
	user := []User{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	json.Unmarshal(body, &user)

	fmt.Println("Response:", user)
	return user, nil
}
