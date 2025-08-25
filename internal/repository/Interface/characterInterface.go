package repository

import "github.com/fbomrl/animon-go/internal/model"

type characterRepositoryInterface interface {
	CharacterById(id int) (*model.Character, error)
}
