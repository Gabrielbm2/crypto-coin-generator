package main

import (
	"database/sql"
	"desafioKlever/models"
	"desafioKlever/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=gbm158545 dbname=Vote_crypto sslmode=disable")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	models.InitDB("user=postgres password=gbm158545 dbname=Vote_crypto host=localhost port=5432 sslmode=disable")

	r := mux.NewRouter()
	routes.CarregaRotas(r)

	fmt.Println("Servidor rodando em http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
