package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (a *API) GetAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := a.Storage.GetUserById(id)
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
