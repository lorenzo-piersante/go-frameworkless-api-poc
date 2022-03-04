package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lorenzo-piersante/go-frameworkless-api-poc/api"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	apiPort := flag.String("port", "8080", "service port")
	flag.Parse()

	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		log.Printf("Cannot connect to database: %v", err)
		os.Exit(1)
	}

	s := storage.NewStorage(db)
	a := api.NewAPI(s)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		<-sigs
		log.Println("performing shutdown...")
		if err := a.Shutdown(); err != nil {
			log.Printf("failed to shutdown server: %v", err)
		}
	}()

	log.Printf("service is ready to listen on port: %s", *apiPort)
	if err := a.Start(*apiPort); err != http.ErrServerClosed {
		log.Printf("server failed: %v", err)
		os.Exit(1)
	}
}
