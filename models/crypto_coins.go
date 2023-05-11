package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
)

//Esse package "models" define a estrutura de dados para a entidade Crypto e suas variações, como CryptoPayload e CryptoWithVotesPayload, além de incluir funções para criar a tabela "crypto_coins" no banco de dados.

type Crypto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// DTOS
type CryptoPayload struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CryptoWithVotesPayload struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}
type CryptosWithVotesPayload struct {
	Cryptos []CryptosWithVotesPayload `json:"cryptos"`
}
type CryptosPayload struct {
	Cryptos []Crypto `json:"cryptos"`
}

type CreateCryptoPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CreateCryptoTable() error {
	sql := `
		CREATE TABLE if not exists "crypto_coins" (
		"id" text PRIMARY KEY,
		"name" text NOT NULL,
		"created_at" timestamp NOT NULL DEFAULT NOW()
	);
	`
	fmt.Println("Criando Tabela de cryptos")
	_, err := db.GetDatabase().Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar tabela de cryptos: %s", err)
	} else {
		fmt.Println("Tabela de cryptos criada")
	}
	return nil
}
