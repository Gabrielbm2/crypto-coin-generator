package main

import (
	"desafioKlever/models"
	"desafioKlever/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Esse é o package principal do projeto, que contém a função main() que é executada quando o programa é iniciado. Ele carrega as configurações do arquivo .env, cria as tabelas do banco de dados e inicia o servidor HTTP, definindo o endereço de IP e porta que serão usados para atender às requisições. Além disso, ele também carrega as rotas do projeto utilizando o roteador chi e define o tratamento de erros caso o servidor falhe em iniciar.

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.CreateCryptoTable()
	models.CreateVotesTable()

	r := chi.NewRouter()
	routes.LoadRoutes(r)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	address := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}

	fmt.Printf("Servidor rodando em %v", address)
	log.Fatal(srv.ListenAndServe())

}
