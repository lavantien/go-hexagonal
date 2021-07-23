package rest

import (
	"github.com/gorilla/mux"
	"github.com/lavantien/go-hexagonal/pkg/adding"
	"github.com/lavantien/go-hexagonal/pkg/reading"
)

func InitHandlers(rs reading.Service, as adding.Service) *mux.Router {
	router := mux.NewRouter()

	// Reading
	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	// Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")

	return router
}
