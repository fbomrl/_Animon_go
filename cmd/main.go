package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fbomrl/animon-go/internal/database"
	"github.com/fbomrl/animon-go/internal/repository"
	"github.com/fbomrl/animon-go/internal/routes"
	"github.com/fbomrl/animon-go/internal/services"
	"github.com/joho/godotenv"
)

var servCharacter *services.CharacterService

func main() {
	//CARREGA .ENV
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DSN não definido na variável DB_DSN")
	}

	//CONECTA BANCO DE DADOS SQLSERVER
	db, err := database.SqlServer(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Conectado ao SQL Server com sucesso!")

	//INJEÇÃO DE DEPENDÊNCIAS
	repoCharacter := &repository.CharacterRepository{DB: db}
	servCharacter = &services.CharacterService{RepoCharacter: repoCharacter}

	//REGISTRO DE ROTAS
	mux := routes.Register(servCharacter)

	//INICIA SERVIDOR
	log.Println("Servidor rodando na porta 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
