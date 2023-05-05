package models

import (
	"context"
	"desafioKlever/db"
	"fmt"
	"time"
)

type CreateCrypto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateCryptoCoinsTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS cryptos (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
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
