package api

import (
	"context"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type API struct {
	Storage *storage.Storage
	server  *http.Server
}

func NewAPI(storage *storage.Storage) *API {
	return &API{
		Storage: storage,
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

	router.POST("/register", a.PostAction)

	router.GET("/users/:id", a.GetAction)

	return router
}

func respond(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("failed to respond: %v", err)
	}
}
