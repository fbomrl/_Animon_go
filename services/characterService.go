package services

import (
	"errors"

	"github.com/fbomrl/animon-go/internal/model"
	"github.com/fbomrl/animon-go/internal/repository/interfaces"
)

var (
	ErrCharacterNotFound = errors.New("personagem não encontrado")
)

type CharacterService struct {
	RepoCharacter interfaces.CharacterRepositoryInterface
}

func (s *CharacterService) CharacterByIdService(id int) (*model.Character, error) {
	character, err := s.RepoCharacter.CharacterById(id)
	if err != nil {
		return nil, err
	}
	if character == nil {
		return nil, ErrCharacterNotFound
	}
	return character, nil
}

func (s *CharacterService) FindAllCharactersService() ([]*model.Character, error) {
	characters, err := s.RepoCharacter.FindAllCharacters()
	if err != nil {
		return nil, err
	}
	if len(characters) == 0 {
		return nil, ErrCharacterNotFound
	}
	return characters, nil
}
