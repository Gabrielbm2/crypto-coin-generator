package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gabrielbm2/desafioKlever/models"
	"github.com/Gabrielbm2/desafioKlever/testutils"
)

func TestCreateCryptoCoin(t *testing.T) {
	cryptoCoin := models.CreateCryptoCoinPayload{
		CryptoID: "BTC",
		Amount:   10,
	}

	// Mock the CreateCryptoCoin function to return no errors.
	testutils.MockCreateCryptoCoin(nil)
	defer testutils.ClearMockData()

	body, err := json.Marshal(cryptoCoin)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/crypto-coins", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCryptoCoin)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestGetCryptoCoin(t *testing.T) {
	cryptoCoinID := "1"
	expectedCryptoCoin := &models.CryptoCoin{
		ID:       cryptoCoinID,
		CryptoID: "BTC",
		Amount:   10,
	}

	// Mock the GetCryptoCoin function to return the expectedCryptoCoin.
	testutils.MockGetCryptoCoin(expectedCryptoCoin, nil)
	defer testutils.ClearMockData()

	req, err := http.NewRequest("GET", "/crypto-coins/"+cryptoCoinID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCryptoCoin)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body.
	var response models.CryptoCoin
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.ID != expectedCryptoCoin.ID || response.CryptoID != expectedCryptoCoin.CryptoID ||
		response.Amount != expectedCryptoCoin.Amount {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response, expectedCryptoCoin)
	}
}

func TestGetCryptoCoins(t *testing.T) {
	// Define a mock response from the service layer
	expectedCoins := []*models.CryptoCoin{
		{
			ID:          1,
			CryptoID:    1,
			Name:        "Bitcoin",
			Description: "The first and most well-known cryptocurrency",
			Price:       50000,
		},
		{
			ID:          2,
			CryptoID:    2,
			Name:        "Ethereum",
			Description: "A programmable blockchain",
			Price:       3000,
		},
	}
	mockService := &service.MockCryptoCoinService{
		GetCryptoCoinsFunc: func() ([]*models.CryptoCoin, error) {
			return expectedCoins, nil
		},
	}
	controller := NewCryptoCoinController(mockService)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/crypto-coins", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder and serve the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetCryptoCoins)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	var response []*models.CryptoCoin
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response) != len(expectedCoins) {
		t.Errorf("handler returned unexpected number of coins: got %v, want %v", len(response), len(expectedCoins))
	}

	for i, coin := range response {
		if coin.ID != expectedCoins[i].ID || coin.CryptoID != expectedCoins[i].CryptoID || coin.Name != expectedCoins[i].Name ||
			coin.Description != expectedCoins[i].Description || coin.Price != expectedCoins[i].Price {
			t.Errorf("handler returned unexpected coin at index %d: got %+v, want %+v", i, coin, expectedCoins[i])
		}
	}
}
