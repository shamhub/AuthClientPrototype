package main

import (
	"fmt"

	"github.com/shamhub/codingtest/api"
)

func main() {
	// client code
	response, err := api.Authenticate(api.PasswordBasedAuthenType,
		&api.PasswordBasedClientAuthenData{Username: "user1",
			Password: "passwd",
			Whatever: 001})
	if err != nil {
		fmt.Println(response.ReadResponse())
	}
}
