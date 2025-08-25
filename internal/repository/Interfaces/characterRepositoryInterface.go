package interfaces

import "github.com/fbomrl/animon-go/internal/model"

type CharacterRepositoryInterface interface {
	CharacterById(id int) (*model.Character, error)
	FindAllCharacters() ([]*model.Character, error)
}
