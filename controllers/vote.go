package controllers

import (
	"desafioKlever/responder"
	"desafioKlever/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Like(w http.ResponseWriter, r *http.Request) {

	cryptoID := chi.URLParam(r, "cryptoID")
	err := services.RegisterVote("like", cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Dislike(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")

	err := services.RegisterVote("dislike", cryptoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetVotes(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")

	votes, err := services.GetVotes(cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, votes, http.StatusOK)
}
