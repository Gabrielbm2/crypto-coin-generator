package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
)

type Crypto struct {
	Id   string `json:"id"`
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
