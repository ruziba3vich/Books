package models

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AuthorId int    `json:"author_id"`
}
