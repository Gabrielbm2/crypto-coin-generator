package controllers

import (
	"desafioKlever/responder"
	"desafioKlever/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetCryptos(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")
	cryptoIDaAsInt, err := strconv.Atoi(cryptoID)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	votes, err := services.GetVotes(cryptoIDaAsInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, votes, http.StatusOK)
}
