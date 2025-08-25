package model

type Character struct {
	Id      int    `json: "id"`
	Nome    string `json: "nome"`
	Alias   string `json: "alias"`
	Species string `json: "species"`
}
