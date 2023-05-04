package routes

import (
	"desafioKlever/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r *chi.Mux) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// TODO: source_ip and source_id in headers
	// fetch data from headers and add to database
	// Check all errors
	r.Post("/crypto/{cryptoID}/like", controllers.Like)
	r.Post("/crypto/{cryptoID}/dislike", controllers.Dislike)

	r.Get("/crypto/{cryptoID}", controllers.GetCrypto) // TODO: retornar votos
	r.Get("/crypto/{cryptoID}/votes", controllers.GetVotes)
	r.Get("/crypto", controllers.GetCryptos)

}
