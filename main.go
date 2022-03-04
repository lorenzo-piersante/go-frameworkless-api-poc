package main

import (
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lorenzo-piersante/go-frameworkless-api-poc/api"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
)

func main() {
	apiPort := flag.String("port", "8080", "service port")
	flag.Parse()

	r := redis.NewClient(&redis.Options{Addr: "localhost:" + *apiPort})
	s := storage.NewStorage(r)
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
