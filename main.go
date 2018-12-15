package main

import (
	firebase "firebase-custom-auth/driver"
	http "firebase-custom-auth/http/rest"
	"fmt"

	"google.golang.org/appengine"
)

func main() {
	//Initializing Firebase
	err := firebase.Init()
	if err != nil {
		fmt.Println("Error initializing firebase:", err.Error())
	}

	http.Run()

	appengine.Main()
}
