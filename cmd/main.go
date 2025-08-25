package main

import (
	"html/template"
	"net/http"
)

type Animon struct {
	Nome       string
	Especie    string
	Subespecie string
	Planeta    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	animon := []Animon{
		{Nome: "Metanik", Especie: "Humand", Subespecie: "Dragoid", Planeta: 5},
		{"Fênix", "Dragoid", "Dinosaurid", 8},
		{"Teste", "Teste", "Teste", 2},
	}

	temp.ExecuteTemplate(w, "Index", animon)
}
