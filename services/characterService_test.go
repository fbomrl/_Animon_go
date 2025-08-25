package services

import (
	"testing"

	"github.com/fbomrl/animon-go/internal/model"
)

type FakeRepository struct {
	character *model.Character
	err       error
}

func (f *FakeRepository) CharacterById(id int) (*model.Character, error) {
	return f.character, f.err
}
func (f *FakeRepository) FindAllCharacters() ([]*model.Character, error) {
	return []*model.Character{f.character}, f.err
}

func TestCharacterByIdService(t *testing.T) {
	repo := &FakeRepository{
		character: &model.Character{
			Id:   1,
			Name: "Metanik Seleriem",
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
}
