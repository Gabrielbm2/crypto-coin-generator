package mappers

import "desafioKlever/models"

//Esse package "mappers" possui funções para mapear modelos de criptomoedas para payloads de criptomoedas, permitindo que as informações sejam transferidas para outras camadas ou sistemas de forma mais organizada e adequada.

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
