package controllers

import (
	"desafioKlever/responder"
	"desafioKlever/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ReturnVotes(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")

	return_votes, err := services.ReturnVotes(cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, return_votes, http.StatusOK)
}
