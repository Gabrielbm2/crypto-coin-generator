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

// Essa função é responsável por testar a funcionalidade de busca de uma criptomoeda específica através do seu ID.
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

// Função para testar criação de crypto moedas
func TestCreateCrypto(t *testing.T) {
	// Cria uma nova criptomoeda
	crypto := models.CreateCryptoPayload{
		ID:   "1",
		Name: "Klever",
	}

	jsonCrypto, _ := json.Marshal(crypto)

	reader := bytes.NewReader(jsonCrypto)

	req := httptest.NewRequest("POST", "/crypto", reader)

	w := httptest.NewRecorder()
	CreateCrypto(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusCreated, w.Code)
	}
}

// Essa função de teste verifica que ela esta retornando corretamente uma lista de criptomoedas cadastradas no sistema
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
