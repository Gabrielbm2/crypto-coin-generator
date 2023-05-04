package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Voto struct {
	ID    int    `json:"id"`
	Opcao string `json:"opcao"`
	Votos int    `json:"votos"`
}

type Votos struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

var db *sql.DB

func InitDB(Vote_crypto string) error {
	var err error
	db, err = sql.Open("postgres", Vote_crypto)
	if err != nil {
		return fmt.Errorf("Erro ao conectar ao banco de dados: %s", err)
	}

	err = criarTabelaVotos()
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatalf("Erro ao fechar conexão com o banco de dados: %s", err)
	}
}

func criarTabelaVotos() error {
	sql := `
	CREATE TABLE IF NOT EXISTS votes (
		id SERIAL PRIMARY KEY,
		opcao TEXT NOT NULL,
		votos INT NOT NULL DEFAULT 0,
		vote_timestamp TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		return fmt.Errorf("Erro ao criar tabela de votos: %s", err)
	}

	return nil
}

func RegistrarVoto(opcao string) error {
	sql := `
	INSERT INTO votes (opcao) VALUES ($1);
	`
	_, err := db.Exec(sql, opcao)
	if err != nil {
		return fmt.Errorf("Erro ao registrar voto: %s", err)
	}

	return nil
}

func ObterVotos() (*Votos, error) {
	sqlLikes := `
	SELECT COUNT(*) FROM votes WHERE opcao = 'like';
	`
	sqlDislikes := `
	SELECT COUNT(*) FROM votes WHERE opcao = 'dislike';
	`

	votos := &Votos{}
	err := db.QueryRow(sqlLikes).Scan(&votos.Likes)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter votos: %s", err)
	}

	err = db.QueryRow(sqlDislikes).Scan(&votos.Dislikes)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter votos: %s", err)
	}

	return votos, nil
}
