package repository

import (
	"database/sql"

	"github.com/fbomrl/animon-go/internal/model"
)

type CharacterRepository struct {
	DB *sql.DB
}

func (repo *CharacterRepository) CharacterById(id int) (*model.Character, error) {
	var character model.Character
	err := repo.DB.QueryRow("SELECT * FROM CHARACTER WHERE ID = ?", id).Scan(
		&character.Id,
		&character.Nome,
		&character.Alias,
		&character.Species,
	)
	if err != nil {
		return nil, err
	}
	return &character, nil
}
