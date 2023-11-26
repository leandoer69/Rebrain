package wallet

import (
	"Rebrain/module09/internal/constants"
	"Rebrain/module09/internal/entities"
	wallet2 "Rebrain/module09/internal/repositories/wallet"
	"encoding/json"
	"net/http"
	"strconv"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	name, pass, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "auth header not found", http.StatusUnauthorized)
		return
	}

	if name != constants.AdminName || pass != constants.AdminPass {
		http.Error(w, "wrong login or password", http.StatusUnauthorized)
		return
	}

	ids, ok := r.URL.Query()["ids"]
	if !ok {
		http.Error(w, "ids not found", http.StatusBadRequest)
		return
	}

	var wallets []entities.Wallet
	for _, id := range ids {
		id, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "wrong id parameter", http.StatusInternalServerError)
			return
		}
		wallet := wallet2.GetById(id)
		if wallet != nil {
			wallets = append(wallets, *wallet)
		}
	}

	jsonData, err := json.Marshal(wallets)
	if err != nil {
		http.Error(w, "invalid json marshall", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}
