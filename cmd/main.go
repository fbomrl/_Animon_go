package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/fbomrl/animon-go/internal/database"
	"github.com/joho/godotenv"
)

var temp = template.Must(template.ParseGlob("../templates/*.html"))

func main() {
	// variaveis .env
	err := godotenv.Load()
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

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	animon := "Teste"

	temp.ExecuteTemplate(w, "Index", animon)
}
