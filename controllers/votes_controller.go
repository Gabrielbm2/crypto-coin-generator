package controllers

import (
	"desafioKlever/responder"
	"desafioKlever/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Like registra um voto positivo para uma criptomoeda específica.
//
// swagger:operation POST /crypto/{cryptoID}/like cryptos like
//
// ---
// parameters:
//   - name: cryptoID
//     in: path
//     description: ID da criptomoeda
//     required: true
//     schema:
//     type: integer
//     format: int64
//
// responses:
//
//	"201":
//	  description: Voto registrado com sucesso.
//	"500":
//	  description: Erro interno do servidor.
func Like(w http.ResponseWriter, r *http.Request) {

	cryptoID := chi.URLParam(r, "cryptoID")
	err := services.RegisterVote("like", cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Dislike registra um voto negativo para a criptomoeda com o ID especificado.
//
// swagger:operation POST /crypto/{cryptoID}/dislike cryptos dislikeCrypto
//
// ---
// parameters:
//   - name: cryptoID
//     in: path
//     description: ID da criptomoeda
//     required: true
//     schema:
//     type: integer
//     format: int64
//
// responses:
//
//	"201":
//	  description: Voto negativo registrado com sucesso
//	"500":
//	  description: Erro interno do servidor
func Dislike(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")

	err := services.RegisterVote("dislike", cryptoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetVotes retorna todos os votos da criptomoeda com o ID especificado.
//
// swagger:operation GET /crypto/{cryptoID}/votes cryptos getVotes
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
// responses:
// "200":
// description: Retorna todos os votos da criptomoeda com o ID especificado.
// content:
// application/json:
// schema:
// type: array
// items:
// $ref: "#/components/schemas/Vote"
// "404":
// description: Criptomoeda não encontrada.
// "500":
// description: Erro interno do servidor.
func GetVotes(w http.ResponseWriter, r *http.Request) {
	cryptoID := chi.URLParam(r, "cryptoID")

	votes, err := services.GetCryptoWithVotesByID(cryptoID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responder.JSON(w, r, votes, http.StatusOK)
}
