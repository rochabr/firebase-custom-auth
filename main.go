package main

import (
	firebase "firebase-authenticator/driver"
	http "firebase-authenticator/http"
	"fmt"
)

func main() {
	//Initializing Firebase
	err := firebase.Init()
	if err != nil {
		fmt.Println("Error initializing firebase:", err.Error())
	}

	//Initializes server
	http.Run()
}
