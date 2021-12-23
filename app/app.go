package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mstreet3/banking-auth/domain"
	"github.com/mstreet3/banking-auth/logger"
	"github.com/mstreet3/banking-auth/service"
)

const appPort int64 = 9000

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

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
	router.Path("/ping").HandlerFunc(pong).Methods(http.MethodGet)
	router.Path("/auth/login").HandlerFunc(authHandlers.login).Methods(http.MethodPost)
	router.Path("/auth/register").HandlerFunc(authHandlers.register).Methods(http.MethodPost)

	err := http.ListenAndServe(uri, router)

	if err != nil {
		log.Fatal(err)
	}
}

func pong(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "pong")
}
