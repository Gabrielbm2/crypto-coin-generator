package services

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
	"fmt"
)

func RegisterVote(opcao, cryptoID string) error {
	db := db.GetDatabase()

	sql := `
	INSERT INTO votes (option,crypto_id) VALUES ($1,$2);
	`
	_, err := db.Exec(context.Background(), sql, opcao, cryptoID)
	if err != nil {
		return fmt.Errorf("Erro ao registrar voto: %s", err)
	}

	return nil
}

func GetVotes(cryptoID string) (*models.Votes, error) {
	db := db.GetDatabase()

	sqlLikes := `
	SELECT COUNT(*) FROM votes WHERE option = 'like' and crypto_id = $1;
	`
	sqlDislikes := `
	SELECT COUNT(*) FROM votes WHERE option = 'dislike'  and crypto_id = $1;
	`

	votos := &models.Votes{}
	err := db.QueryRow(context.Background(), sqlLikes, cryptoID).Scan(&votos.Likes)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter votos: %s", err)
	}

	err = db.QueryRow(context.Background(), sqlDislikes, cryptoID).Scan(&votos.Dislikes)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter votos: %s", err)
	}

	return votos, nil
}
