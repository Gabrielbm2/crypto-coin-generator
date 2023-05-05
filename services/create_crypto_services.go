package services

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
	"fmt"

	"github.com/google/uuid"
)

func CreateCrypto(newCrypto NewCrypto) (models.Crypto, error) {
	id := uuid.New().String()

	_, err := db.GetDatabase().Exec(context.Background(), `
        INSERT INTO crypto_coins (id, name) VALUES ($1, $2)
    `, id, newCrypto.Name)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar nova moeda: %s", err)
	}

	crypto := &CreateCrypto{
		Id:   id,
		Name: newCrypto.Name,
	}

	return crypto, nil
}
