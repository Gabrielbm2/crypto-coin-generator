package routes

import (
	"desafioKlever/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//O package "routes" define as rotas (endpoints) da aplicação, mapeando cada uma para a função correspondente em "controllers" e incluindo também uma rota para servir arquivos estáticos.

func LoadRoutes(r *chi.Mux) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/crypto", controllers.GetCryptos)
	r.Post("/crypto", controllers.CreateCrypto)
	r.Get("/crypto/{cryptoID}", controllers.GetCrypto)

	r.Get("/crypto/{cryptoID}/votes", controllers.GetVotes)
	r.Post("/crypto/{cryptoID}/like", controllers.Like)
	r.Post("/crypto/{cryptoID}/dislike", controllers.Dislike)

	// Docs
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

}
