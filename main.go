package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

	mux := chi.NewRouter()
	mux.Get("/status", handleStatus)

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%s", port),
		Handler:   mux,
		TLSConfig: nil,
	}

	err = srv.ListenAndServe()
}
