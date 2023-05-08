package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetCrypto(t *testing.T) {
	req := httptest.NewRequest("GET", "/crypto/1", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("cryptoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	w := httptest.NewRecorder()
	GetCrypto(w, req)

	expectedResponse := `{"id":"1","name":"Bitcoin"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta incorreta. Esperado: %s. Obtido: %s", expectedResponse, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusOK, w.Code)
	}
}
