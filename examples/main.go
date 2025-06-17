package main

import (
	"github.com/TheGP/vrchat-go-with-proxy"
)

func main() {
	client := vrchat.NewClient("https://api.vrchat.cloud/api/1")

	err := client.Authenticate("username", "password", "totp", "user-agent")
	if err != nil {
		panic(err)
	}

	user, err := client.GetCurrentUser()
	if err != nil {
		panic(err)
	}

	println("logged in as ", user.DisplayName)
}
