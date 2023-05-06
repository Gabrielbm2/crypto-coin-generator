package mappers

import "desafioKlever/models"

func MapCryptosToPayload(cryptos []*models.Crypto) []*models.CryptoPayload {
	cryptosPayload := make([]*models.CryptoPayload, len(cryptos))
	for i, crypto := range cryptos {
		cryptosPayload[i] = MapCryptoToPayload(crypto)
	}
	return cryptosPayload
}

func MapCryptoToPayload(crypto *models.Crypto) *models.CryptoPayload {
	return &models.CryptoPayload{
		Id:   crypto.Id,
		Name: crypto.Name,
	}

}
