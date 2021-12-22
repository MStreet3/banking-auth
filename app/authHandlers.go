package app

import (
	"encoding/json"
	"net/http"

	"github.com/mstreet3/banking-auth/dto"
	"github.com/mstreet3/banking-auth/service"
)

type AuthHandlers struct {
	service service.AuthService
}

func (h AuthHandlers) register(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusNotImplemented, "Handler not implemented..")
}

func (h AuthHandlers) login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	/* get access token */
	token, appErr := h.service.Login(loginRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, token)

}
