package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstreet3/banking-auth/logger"
)

const appPort int64 = 9000

func pong(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "pong")
}

func Start() {
	logger.Info(fmt.Sprintf("Starting the application on port %d", appPort))
	uri := fmt.Sprintf("localhost:%d", appPort)

	router := mux.NewRouter()
	router.Path("/ping").HandlerFunc(pong).Methods(http.MethodGet)

	err := http.ListenAndServe(uri, router)

	if err != nil {
		log.Fatal(err)
	}
}
