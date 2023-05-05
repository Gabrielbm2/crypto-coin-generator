package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
)

type ReturnVotes struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

func ReturnVotesTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS votes (
		id SERIAL PRIMARY KEY,
		crypto_id VARCHAR NOT NULL,
		vote_type VARCHAR NOT NULL
);`

	fmt.Println("Criando Tabela de votes")
	_, err := db.GetDatabase().Exec(context.Background(), sql)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar tabela de votos: %s", err)
	} else {
		fmt.Println("Tabela de votos criada")
	}

	return nil
}
