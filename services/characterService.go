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
	repoCharacter interfaces.CharacterRepositoryInterface
}

func (s *CharacterService) CharacterByIdService(id int) (*model.Character, error) {
	character, err := s.repoCharacter.CharacterById(id)
	if err != nil {
		return nil, err
	}
	if character == nil {
		return nil, ErrCharacterNotFound
	}
	return character, nil
}

func (s *CharacterService) FindAllCharactersService() ([]*model.Character, error) {
	characters, err := s.repoCharacter.FindAllCharacters()
	if err != nil {
		return nil, err
	}
	if len(characters) == 0 {
		return nil, ErrCharacterNotFound
	}
	return characters, nil
}
