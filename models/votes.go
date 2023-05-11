package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
)

//Esse package models define as estruturas de dados que são usadas no sistema, incluindo Crypto e Votes. Ele também inclui algumas estruturas DTO (Data Transfer Objects) para representar esses dados em diferentes contextos, como CryptoPayload e VotesPayload. Além disso, o pacote também contém funções para criar tabelas no banco de dados, como CreateCryptoTable e CreateVotesTable.

type Votes struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type VoteInfo struct {
	VoteType   int    `json:"voteType"`
	CryptoID   int    `json:"cryptoID"`
	CryptoName string `json:"cryptoName"`
}

// DTOS
type VotesPayload struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

func CreateVotesTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS "votes" (
  	"id" serial PRIMARY KEY,
  	"option" text NOT NULL,
  	"crypto_id" text references crypto_coins(id) NOT NULL,
  	"source_ip" text,
  	"source_id" text UNIQUE,
  	"created_at" timestamp NOT NULL DEFAULT NOW()
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
