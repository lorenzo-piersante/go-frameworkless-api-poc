package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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
		Handler: a.bootRouter,
	}

	return a.server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

func (a *API) bootRouter() *httprouter.Router {
	router := httprouter.New()

	// rotte
	router.GET("/users/:id", a.Get)

	return router
}

func (a *API) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("failed to decode body: %v", err)
		write(w, 400, nil)
		return
	}

	err := a.storage.UpdateUserById(id, &user)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		write(w, 500, nil)
		return
	}

	write(w, 200, okResponse())
}

func okResponse() []byte {
	return []byte(`{"message":"ok"}`)
}

func write(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("failed to write: %v", err)
	}
}
