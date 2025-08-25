package repository

import (
	"database/sql"

	"github.com/fbomrl/animon-go/internal/character/model"
)

type CharacterRepository struct {
	DB *sql.DB
}

func (repo *CharacterRepository) CharacterById(id int) (*model.Character, error) {
	var character model.Character
	err := repo.DB.QueryRow("SELECT * FROM CHARACTER WHERE ID = ?", id).Scan(
		&character.Id,
		&character.Name,
		&character.Alias,
		&character.Species,
	)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (repo *CharacterRepository) FindAllCharacters() ([]*model.Character, error) {
	rows, err := repo.DB.Query("SELECT ID, NAME, ALIAS, SPECIES FROM CHARACTER")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*model.Character

	for rows.Next() {
		var character model.Character
		err := rows.Scan(&character.Id, &character.Name, &character.Alias, &character.Species)
		if err != nil {
			return nil, err
		}
		characters = append(characters, &character)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}
