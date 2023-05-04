package routes

import (
	"desafioKlever/controllers"

	"github.com/gorilla/mux"
)

func CarregaRotas(r *mux.Router) {
	r.HandleFunc("/like", controllers.Like)
	r.HandleFunc("/dislike", controllers.Dislike)
}
