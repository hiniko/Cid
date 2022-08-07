package main

import "net/http"

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}
