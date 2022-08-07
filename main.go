package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
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

	srv := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: nil,
	}

	err = srv.ListenAndServe()
}
