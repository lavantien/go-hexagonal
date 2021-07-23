package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lavantien/go-hexagonal/pkg/adding"
	"github.com/lavantien/go-hexagonal/pkg/http/rest"
	"github.com/lavantien/go-hexagonal/pkg/reading"
	"github.com/lavantien/go-hexagonal/pkg/storage"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("Error while setting up storage:", err)
	}

	rs := reading.NewService(r)
	as := adding.NewService(r)

	fmt.Println("Starting server on port 8080...")
	router := rest.InitHandlers(rs, as)
	log.Fatal(http.ListenAndServe(":8080", router))
}
