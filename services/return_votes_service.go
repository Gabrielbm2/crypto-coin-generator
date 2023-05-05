package services

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
)

func ReturnVotes(cryptoID string) (*models.Crypto, error) {
	var likes, dislikes int

	err := db.GetDatabase().QueryRow(context.Background(),
		"SELECT COUNT(*) FROM votes WHERE crypto_id = $1 AND vote_type = 'like'",
		cryptoID).Scan(&likes)
	if err != nil {
		return nil, err
	}

	err = db.GetDatabase().QueryRow(context.Background(),
		"SELECT COUNT(*) FROM votes WHERE crypto_id = $1 AND vote_type = 'dislike'",
		cryptoID).Scan(&dislikes)
	if err != nil {
		return nil, err
	}

	return likes, dislikes, nil
}
