package controllers

import (
	"desafioKlever/responder"
	"desafioKlever/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetCrypto(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")
	crypto, err := services.GetCrypto(cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, crypto, http.StatusOK)
}
