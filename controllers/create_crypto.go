package controllers

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func CreateCrypto(w http.ResponseWriter, r *http.Request) {
	var crypto models.Crypto
	err := json.NewDecoder(r.Body).Decode(&crypto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Erro ao decodificar dados da criptomoeda:", err)
		return
	}

	conn := db.GetDatabase()
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao iniciar transação no banco de dados:", err)
		return
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO cryptos (name, symbol, price) VALUES ($1, $2)", crypto.Name, crypto.Id)
	if err != nil {
		tx.Rollback(context.Background())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao inserir criptomoeda no banco de dados:", err)
		return
	}

	err = tx.Commit(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erro ao finalizar transação no banco de dados:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
