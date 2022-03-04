package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"

	"github.com/julienschmidt/httprouter"
)

type API struct {
	storage *storage.Storage
	server  *http.Server
}

func NewAPI(storage *storage.Storage) *API {
	return &API{
		storage: storage,
	}
}

func (a *API) Start(port string) error {
	a.server = &http.Server{
		Addr:    ":" + port,
		Handler: a.bootRouter(),
	}

	return a.server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

func (a *API) bootRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/users/:id", a.GetAction)
	router.POST("/users", a.PostAction)

	return router
}

func (a *API) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func respond(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("failed to respond: %v", err)
	}
}
