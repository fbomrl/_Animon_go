package services

import (
	"fmt"
	"testing"

	"github.com/fbomrl/animon-go/internal/model"
)

type FakeRepository struct {
	characters []*model.Character
	err        error
}

func (f *FakeRepository) CharacterById(id int) (*model.Character, error) {

	for _, c := range f.characters {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, fmt.Errorf("personagem não encontrado")
}
func (f *FakeRepository) FindAllCharacters() ([]*model.Character, error) {
	if len(f.characters) > 0 {
		return f.characters, nil
	}
	return nil, fmt.Errorf("nenhum personagem encontrado")
}

func TestCharacterByIdService(t *testing.T) {
	//PERSONAGEM EXISTE
	repo := &FakeRepository{
		characters: []*model.Character{
			&model.Character{Id: 1, Name: "Metanik Seleriem"},
		},
		err: nil,
	}
	characterService := &CharacterService{RepoCharacter: repo}
	character, err := characterService.CharacterByIdService(1)
	if err != nil {
		t.Errorf("Teste: esperado sem erro - recebido: %v", err)
		return
	}
	if character == nil {
		t.Errorf("Teste: esperado personagem não nulo - recebido: nil")
	}

	//PERSONAGEM NÃO EXISTE
	character, err = characterService.CharacterByIdService(-1)
	if err == nil {
		t.Errorf("Teste: esperado erro - recebido: nil")
	}
	if character != nil {
		t.Errorf("Teste: esperado nil - recebido: %v", character)
	}
}

func TestFindAllCharacters(t *testing.T) {
	repo := &FakeRepository{
		characters: []*model.Character{
			{Id: 1, Name: "Metanik Seleriem", Alias: "Metanik", Species: "1"},
			{Id: 8, Name: "Belchior", Alias: "Theoty", Species: "1"},
		},
		err: nil,
	}
	characterService := &CharacterService{RepoCharacter: repo}

	allCharacters, err := characterService.FindAllCharactersService()
	if err != nil {
		t.Errorf("Teste: esperado sem erro - recebido: %v", err)
		return
	}
	if allCharacters == nil {
		t.Errorf("Teste: lista personagens não nulo - recebido: nil")
	}
}
