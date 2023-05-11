package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Função que testa like
func TestLike(t *testing.T) {
	req := httptest.NewRequest("POST", "/crypto/1/vote/like", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("cryptoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	w := httptest.NewRecorder()
	Like(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusCreated, w.Code)
	}
}

// Função que testa dislike
func TestDislike(t *testing.T) {
	req := httptest.NewRequest("POST", "/crypto/1/dislike", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("cryptoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	w := httptest.NewRecorder()
	Dislike(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusCreated, w.Code)
	}
}

// Função que testa a quantidade de votos de uma crypto
func TestGetVotes(t *testing.T) {
	req := httptest.NewRequest("GET", "/votes/1", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("cryptoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	w := httptest.NewRecorder()
	GetVotes(w, req)

	expectedResponse := `{"id":"1","name":"Klever","votes":10}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta incorreta. Esperado: %s. Obtido: %s", expectedResponse, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Errorf("Status code incorreto. Esperado: %d. Obtido: %d", http.StatusOK, w.Code)
	}
}
