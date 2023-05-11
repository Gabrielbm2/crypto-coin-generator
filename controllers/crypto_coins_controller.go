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

// GetCrypto recupera as informações de uma criptomoeda específica pelo seu ID.
//
// swagger:operation GET /crypto/{cryptoID} cryptos getCryptoByID
//
// ---
// parameters:
// - name: cryptoID
// in: path
// description: ID da criptomoeda
// required: true
// schema:
// type: integer
// format: int64
//
// responses:
// "200":
// description: Retorna as informações da criptomoeda
// schema:
// "$ref": "#/definitions/Crypto"
// "404":
// description: Criptomoeda não encontrada
// "500":
// description: Erro interno do servidor
func GetCrypto(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")
	crypto, err := services.GetCrypto(cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, crypto, http.StatusOK)
}

// CreateCrypto cria uma nova criptomoeda.
//
// swagger:operation POST /crypto cryptos createCrypto
//
// ---
// parameters:
// - name: payload
// in: body
// description: Dados da criptomoeda a ser criada
// required: true
// schema:
// "$ref": "#/definitions/CreateCryptoPayload"
//
// responses:
//
// "201":
// description: Criptomoeda criada com sucesso
// "400":
// description: Erro ao decodificar dados da criptomoeda
// "500":
// description: Erro interno do servidor ao finalizar transação no banco de dados
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

// GetCryptos retorna todas as criptomoedas cadastradas no sistema.
//
// swagger:operation GET /cryptos cryptos getCryptos
//
// ---
// responses:
// "200":
// description: Retorna um array com todas as criptomoedas cadastradas.
// schema:
// type: array
// items:
// $ref: "#/definitions/CryptoPayload"
// "500":
// description: Erro interno do servidor.
func GetCryptos(w http.ResponseWriter, r *http.Request) {

	cryptos, err := services.GetAllCryptos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao finalizar transação no banco de dados:", err)
		return
	}

	responder.JSON(w, r, mappers.MapCryptosToPayload(cryptos), http.StatusOK)
}
