package models

type Book struct {
	BookId     string `json:"bookId,omitempty" bson:"bookId,omitempty"`
	Name       string `json:"name,omitempty" bson:"name"`
	Autor      string `json:"autor,omitempty" bson:"autor"`
	Genero     string `json:"genero,omitempty" bson:"genero"`
	Preco      string `json:"preco,omitempty" bson:"preco"`
	Prioridade string `json:"prioridade,omitempty" bson:"prioridade"`
}
