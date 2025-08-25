package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/fbomrl/animon-go/services"
)

func CharacterByIdHandler(s *services.CharacterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		idString := strings.TrimPrefix(url, "/characters/id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Id inválido", http.StatusNotFound)
			return
		}

		characterById, err := s.CharacterByIdService(id)
		if err != nil {
			http.Error(w, "Erro ao encontrar personagem", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(characterById)
	}
}
func FindAllCharactersHandler(s *services.CharacterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		characters, err := s.FindAllCharactersService()

		if err != nil {
			http.Error(w, "Erro ao buscar personagens", http.StatusInternalServerError)
			return
		}

		allCharacters, err := json.MarshalIndent(characters, "", " ")
		if err != nil {
			http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		}
		w.Write(allCharacters)
	}
}
