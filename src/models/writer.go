package models

type Writer struct {
	WriterId  string `json:"writerId,omitempty" bson:"writerId,omitempty"`
	Telefone  string `json:"telefone,omitempty" bson:"-"`
	Name      string `json:"name,omitempty" bson:"-"`
	Email     string `json:"email,omitempty" bson:"-"`
	PassWord  string `json:"password,omitempty" bson:"-"`
	Rg        string `json:"rg,omitempty" bson:"-"`
	Cpf       string `json:"cpf,omitempty" bson:"-"`
	City      string `json:"city,omitempty" bson:"-"`
	BirthDate string `json:"birthDate,omitempty" bson:"-"`
	Image     string `json:"image,omitempty" bson:"-"`
}
