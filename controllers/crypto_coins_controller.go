package controllers

import (
	"desafioKlever/mappers"
	"desafioKlever/models"
	"desafioKlever/responder"
	"desafioKlever/services"
	"encoding/json"
	"fmt"
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

func CreateCrypto(w http.ResponseWriter, r *http.Request) {
	var crypto models.CreateCryptoPayload
	err := json.NewDecoder(r.Body).Decode(&crypto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao decodificar dados da criptomoeda:", err)
		return
	}

	_, err = services.CreateCrypto(crypto.ID, crypto.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao finalizar transação no banco de dados:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCryptos(w http.ResponseWriter, r *http.Request) {

	cryptos, err := services.GetAllCryptos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao finalizar transação no banco de dados:", err)
		return
	}

	responder.JSON(w, r, mappers.MapCryptosToPayload(cryptos), http.StatusOK)
}
