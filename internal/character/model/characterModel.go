package model

type Character struct {
	Id      int    `json: "id"`
	Name    string `json: "name"`
	Alias   string `json: "alias"`
	Species string `json: "species"`
}
