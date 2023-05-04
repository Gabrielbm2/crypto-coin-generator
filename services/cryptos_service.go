package services

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
	"fmt"
)

func GetCrypto(cryptoID string) (*models.Crypto, error) {
	db := db.GetDatabase()

	sqlCrypto := `
	SELECT id,name FROM crypto_coins WHERE id = $1;
	`

	crypto := &models.Crypto{}
	err := db.QueryRow(context.Background(), sqlCrypto, cryptoID).Scan(&crypto.Id, &crypto.Name)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao obter crypto: %s", err)
	}

	return crypto, nil
}
