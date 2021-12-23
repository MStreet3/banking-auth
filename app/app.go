package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstreet3/banking-auth/domain"
	"github.com/mstreet3/banking-auth/logger"
	"github.com/mstreet3/banking-auth/service"
)

/* todo: move to environment variables */
const appPort int64 = 9000

func Start() {
	/* todo: add middleware logging of requests and responses */
	logger.Info(fmt.Sprintf("Starting the application on port %d", appPort))
	uri := fmt.Sprintf("localhost:%d", appPort)

	router := mux.NewRouter()

	/* define data source */
	dbClient := getDbClient()
	authRepository := domain.NewAuthRepositoryDb(dbClient)
	authService := service.NewAuthService(authRepository)
	authHandlers := AuthHandlers{
		service: authService,
	}

	/* setup routes to handle */
	router.Path("/ping").
		HandlerFunc(pong).
		Methods(http.MethodGet)
	router.Path("/auth/login").
		HandlerFunc(authHandlers.login).
		Methods(http.MethodPost)
	router.Path("/auth/verify").
		HandlerFunc(authHandlers.verify).
		Methods(http.MethodGet).
		Queries("token", "{token}")
	router.Path("/auth/register").
		HandlerFunc(authHandlers.register).
		Methods(http.MethodPost)

	err := http.ListenAndServe(uri, router)

	if err != nil {
		log.Fatal(err)
	}
}
