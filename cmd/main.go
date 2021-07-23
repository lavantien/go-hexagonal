package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lavantien/go-hexagonal/pkg/http/rest"
	"github.com/lavantien/go-hexagonal/pkg/storage"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("Error while setting up storage:", err)
	}

	c, err := r.GetAllCandyNames()
	if err != nil {
		log.Fatalln("Error while getting candies in storage:", err)
	}
	log.Println(c)

	fmt.Println("Starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}
