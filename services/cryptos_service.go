package services

import (
	"context"
	"desafioKlever/db"
	"desafioKlever/models"
	"fmt"
)

//Esse package "services" define funções para interagir com a camada de dados relacionada às criptomoedas, incluindo GetCrypto para obter informações sobre uma criptomoeda por ID, CreateCrypto para criar uma nova criptomoeda e GetAllCryptos para obter todas as criptomoedas armazenadas no banco de dados. As funções usam o pacote db para se comunicar com o banco de dados.

func GetCrypto(id string) (*models.Crypto, error) {
	db := db.GetDatabase()

	sqlCrypto := `
	SELECT id,name FROM crypto_coins WHERE id = $1;
	`

	crypto := &models.Crypto{}
	err := db.QueryRow(context.Background(), sqlCrypto, id).Scan(&crypto.Id, &crypto.Name)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao obter crypto: %s", err)
	}

	return crypto, nil
}

func CreateCrypto(id, name string) (*models.Crypto, error) {
	db := db.GetDatabase()

	sqlCrypto := `
        INSERT INTO crypto_coins (id, name) VALUES ($1, $2)
    `
	newCrypto := &models.Crypto{}
	_, err := db.Exec(context.Background(), sqlCrypto, id, name)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao obter crypto: %s", err)
	}

	return newCrypto, nil
}

func GetAllCryptos() ([]*models.Crypto, error) {
	db := db.GetDatabase()

	cryptos := []*models.Crypto{}

	sqlCrypto := `
        select id,name from crypto_coins;
    `
	rows, err := db.Query(context.Background(), sqlCrypto)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("erro ao obter cryptos: %s", err)
	}
	for rows.Next() {
		crypto := models.Crypto{}
		rows.Scan(&crypto.Id, &crypto.Name)
		cryptos = append(cryptos, &crypto)
	}

	return cryptos, nil
}
