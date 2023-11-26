package gorilla

import (
	"Rebrain/module09/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", user.CreateUserHandler).Methods(http.MethodPost)
	userRouter.HandleFunc("/", user.GetUserListHandler).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9]+}", user.GetUserHandler).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9]+}", user.DeleteUserHandler).Methods(http.MethodDelete)

	http.ListenAndServe("127.0.0.1:8080", r)
	//http.HandleFunc("/", user.GetUserHandler)
	//http.HandleFunc("/", user.GetUserListHandler)
	//http.HandleFunc("/", user.DeleteUserHandler)
}
