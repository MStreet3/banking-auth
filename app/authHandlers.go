package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstreet3/banking-auth/dto"
	"github.com/mstreet3/banking-auth/logger"
	"github.com/mstreet3/banking-auth/service"
)

type AuthHandlers struct {
	service service.AuthService
}

func (h AuthHandlers) register(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusNotImplemented, "Handler not implemented..")
}

func (h AuthHandlers) verify(w http.ResponseWriter, r *http.Request) {

	/* fetch URL query params */
	vars := mux.Vars(r)
	token := vars["token"]

	/* validate the query params */
	if token == "" {
		writeResponse(w, http.StatusBadRequest, "invalid access token")
	}

	/* attempt to parse JWT for claims */
	claims, appErr := h.service.ParseClaims(token)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Error())
		return
	}

	/* return response */
	logger.Info(fmt.Sprintf("returning claims: %v", claims))
	writeResponse(w, http.StatusOK, claims)
}

func (h AuthHandlers) login(w http.ResponseWriter, r *http.Request) {

	/* attempt to deserialize the request */
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	/* validate the request */
	/* todo: move to middleware */
	invalid := loginRequest.Validate()
	if invalid != nil {
		writeResponse(w, http.StatusBadRequest, invalid.Error())
		return
	}

	/* attempt to get access token */
	token, appErr := h.service.Login(loginRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Error())
		return
	}

	/* return response */
	writeResponse(w, http.StatusOK, token)

}
