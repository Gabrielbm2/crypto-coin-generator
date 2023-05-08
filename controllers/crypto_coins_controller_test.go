package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Gabrielbm2/desafioKlever/models"
	"github.com/Gabrielbm2/desafioKlever/services/mocks"
)

func TestGetCryptoByID(t *testing.T) {
	// Cria um serviço mock
	mockService := new(mocks.CryptoService)

	// Cria um controlador com o serviço mock injetado
	controller := NewCryptoController(mockService)

	// Cria uma criptomoeda mock
	crypto := &models.Crypto{
		ID:    1,
		Name:  "Bitcoin",
		Price: 45000.00,
	}

	// Define o ID de criptomoeda a ser usado na requisição
	cryptoID := "1"

	// Define o comportamento esperado do serviço mock quando a função GetCryptoByID for chamada
	mockService.On("GetCryptoByID", cryptoID).Return(crypto, nil)

	// Cria uma requisição HTTP GET
	req, err := http.NewRequest("GET", "/crypto-coins/"+cryptoID, nil)
	assert.NoError(t, err)

	// Cria um registrador de respostas HTTP e faz a requisição
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetCryptoByID)
	handler.ServeHTTP(rr, req)

	// Verifica se o status da resposta é OK (200)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verifica se a resposta tem o formato esperado
	expectedResponse := `{"ID":1,"Name":"Bitcoin","Price":45000}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verifica se a função GetCryptoByID do serviço mock foi chamada com os argumentos corretos
	mockService.AssertCalled(t, "GetCryptoByID", cryptoID)
}

func NewCryptoController(mockService invalid type) {
	panic("unimplemented")
}
