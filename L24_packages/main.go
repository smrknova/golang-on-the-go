package main

import (
	"fmt"
	"github.com/smrknova/podcast/auth"
	"github.com/smrknova/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("nova", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}