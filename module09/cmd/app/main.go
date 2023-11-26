package main

import (
	"Rebrain/module09/internal/handlers/wallet"
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
	Message string
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.Message)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/source/", http.StripPrefix("/source/",
		http.FileServer(http.Dir("C:\\Users\\Глеб\\GolandProjects\\Rebrain\\module09\\cmd\\app"))))
	http.Handle("/longPing", http.TimeoutHandler(MyHandler{Message: "You're goddamn right"},
		time.Second, "Request timeout"))
	http.HandleFunc("/wallet", wallet.Handle)

	//	gorilla.Start()
	// chiRouter.Start()
	server := http.Server{Addr: "127.0.0.1:8080"}
	server.ListenAndServe()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello human")
}
