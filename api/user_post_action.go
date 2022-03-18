package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/core"
	"net/http"
)

type PostActionInput struct {
	Username string
	Password string
}

type PostActionOutput struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (a *API) PostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode request
	decoder := json.NewDecoder(r.Body)
	var input PostActionInput
	err := decoder.Decode(&input)
	if err != nil {
		respond(w, 400, []byte(`{"message":"failed decoding request body"}`))
		return
	}

	// delegate every logic to a "core" service
	user, err := core.CreateUser(a.Storage, input.Username, input.Password)
	if err != nil || user == nil {
		respond(w, 500, []byte(`{"message":"internal server error"}`))
		return
	}

	// encode response
	encoder := json.NewEncoder(w)
	err = encoder.Encode(PostActionOutput{user.Id, user.Username})
	if err != nil {
		respond(w, 500, []byte(`{"message":"failed encoding response"}`))
		return
	}

	respond(w, 200, []byte(""))
}
