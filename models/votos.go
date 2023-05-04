package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
)

type Votes struct {
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
