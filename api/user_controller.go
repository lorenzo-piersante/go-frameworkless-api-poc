package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (a *API) GetAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := a.storage.GetUserById(id)
	if err != nil {
		respond(w, 500, []byte(`{"message":"internal server error"}`))
		return
	}

	if user == nil {
		respond(w, 404, []byte(`{"message":"user not found"}`))
		return
	}

	encodedUser, err := json.Marshal(user)
	if err != nil {
		respond(w, 500, []byte(`{"message":"internal server error"}`))
		return
	}

	respond(w, 200, encodedUser)
	return
}

type PostActionInput struct {
	username string
	password string
}

type PostActionOutput struct {
	id       string
	username string
	password string
}

func (a *API) PostAction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var input PostActionInput
	err := decoder.Decode(&input)
	if err != nil {
		respond(w, 400, []byte(`{"message":"failed decoding request body"}`))
		return
	}

	user, err := a.storage.CreateUser(input.username, input.password)
	if err != nil || user == nil {
		respond(w, 500, []byte(`{"message":"internal server error"}`))
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(user)
	if err != nil {
		respond(w, 500, []byte(`{"message":"failed encoding response"}`))
		return
	}

	respond(w, 200, []byte(""))
}
