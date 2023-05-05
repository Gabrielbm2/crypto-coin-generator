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

	r.Post("/crypto/{cryptoID}/like", controllers.Like)
	r.Post("/crypto/{cryptoID}/dislike", controllers.Dislike)

	r.Get("/crypto/{cryptoID}", controllers.ReturnVotes)
	r.Get("/crypto/{cryptoID}/votes", controllers.GetVotes)
	r.Get("/crypto", controllers.CreateCrypto)

}
