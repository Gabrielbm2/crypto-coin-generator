package controllers

import (
	"bytes"
	"context"
	"desafioKlever/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// TEST
func TestGetCrypto(t *testing.T) {
	req := httptest.NewRequest("GET", "/crypto/1", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("cryptoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	w := httptest.NewRecorder()
	GetCrypto(w, req)

	expectedResponse := `{"id":"1","name":"Klever"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta incorreta. Esperado: %s. Obtido: %s", expectedResponse, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusOK, w.Code)
	}
}

func TestCreateCrypto(t *testing.T) {
	// Cria uma nova criptomoeda
	crypto := models.CreateCryptoPayload{
		ID:   "1",
		Name: "Klever",
	}

	// Codifica a criptomoeda em JSON
	jsonCrypto, _ := json.Marshal(crypto)

	// Cria um reader com os dados codificados em JSON
	reader := bytes.NewReader(jsonCrypto)

	// Cria um novo request com o reader como corpo
	req := httptest.NewRequest("POST", "/crypto", reader)

	w := httptest.NewRecorder()
	CreateCrypto(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusCreated, w.Code)
	}
}

func TestGetCryptos(t *testing.T) {
	req := httptest.NewRequest("GET", "/cryptos", nil)
	w := httptest.NewRecorder()

	GetCryptos(w, req)

	expectedResponse := `[{"id":"1","name":"Bitcoin"},{"id":"2","name":"Ethereum"},{"id":"3","name":"Binance Coin"}]`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta incorreta. Esperado: %s. Obtido: %s", expectedResponse, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusOK, w.Code)
	}
}
