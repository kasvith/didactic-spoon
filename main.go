package main

import (
	"github.com/kasvith/didactic-spoon/nativestore"
	"log"
)

func main() {
	err := nativestore.Set("example-key", "https://com.example.g", "admin", "damin")
	if err != nil {
		log.Fatal(err)
	}

	usr, secret, err := nativestore.Get("example-key", "https://com.example.g")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(usr, secret)
}
