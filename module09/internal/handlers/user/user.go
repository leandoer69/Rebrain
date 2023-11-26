package user

import (
	"Rebrain/module09/internal/entities"
	userRep "Rebrain/module09/internal/repositories/user"
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body failed", http.StatusBadRequest)
		return
	}

	var user entities.User
	err = json.Unmarshal(jsonBody, &user)
	if err != nil {
		http.Error(w, "failed unmarshalling json", http.StatusBadRequest)
		return
	}

	err = userRep.Create(user)
	if err != nil {
		http.Error(w, "failed creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)		// gorilla
	//id, _ := strconv.Atoi(vars["id"])

	id, _ := strconv.Atoi(chi.URLParam(r, "id")) // chi

	user, err := userRep.GetUser(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "failed marshalling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	users := userRep.GetList()
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "failed marshalling json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)	// for gorilla
	//id, _ := strconv.Atoi(vars["id"])
	id, _ := strconv.Atoi(chi.URLParam(r, "id")) // for chi

	err := userRep.Delete(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
