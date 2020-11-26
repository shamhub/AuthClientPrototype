package main

import (
	"fmt"

	"github.com/shamhub/codingtest/cmd"
)

func main() {
	// client code
	response, err := cmd.Authenticate(cmd.PasswordBasedAuthenType, &cmd.PasswordBasedClientAuthenData{"user1", "passwd", 001})
	if err != nil {
		fmt.Println("not nil")
		fmt.Println(response.ReadResponse())
	} else {
		fmt.Println("nil")
		fmt.Println(response.ReadResponse())
	}

}
