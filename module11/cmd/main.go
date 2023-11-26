package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	mux := chi.NewMux()
	mux.HandleFunc("/hello", helloHandler)

	_ = http.ListenAndServe("localhost:8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "YOU'RE GODDAMN RIGHT!")
}
