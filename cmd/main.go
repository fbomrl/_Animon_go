package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/fbomrl/animon-go/internal/database"
	"github.com/fbomrl/animon-go/internal/handlers"
	"github.com/fbomrl/animon-go/internal/repository"
	"github.com/fbomrl/animon-go/internal/services"
	"github.com/joho/godotenv"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))
var servCharacter *services.CharacterService

func methodHandler(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}

func main() {
	// Carrega .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DSN não definido na variável DB_DSN")
	}

	db, err := database.SqlServer(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conectado ao SQL Server com sucesso!")

	// Inicializa repositório e serviço
	repoCharacter := &repository.CharacterRepository{DB: db}
	servCharacter = &services.CharacterService{RepoCharacter: repoCharacter}

	// Rotas
	http.HandleFunc("/", index)
	http.HandleFunc("/characters/id", methodHandler("GET", handlers.CharacterByIdHandler(servCharacter)))
	http.HandleFunc("/characters", handlers.FindAllCharactersHandler(servCharacter))

	// Inicia servidor
	log.Println("Servidor rodando na porta 8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	characters, err := servCharacter.FindAllCharactersService()
	if err != nil {
		http.Error(w, "Erro ao buscar personagens", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "Index", characters)
}
