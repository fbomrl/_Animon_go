package routes

import (
	"net/http"

	"github.com/fbomrl/animon-go/internal/handlers"
	"github.com/fbomrl/animon-go/internal/services"
)

func Register(service *services.CharacterService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index(service))
	mux.HandleFunc("/characters/id", handlers.CharacterByIdHandler(service))
	mux.HandleFunc("/characters", handlers.FindAllCharactersHandler(service))

	return mux
}
