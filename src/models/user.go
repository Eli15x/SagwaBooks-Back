package models

type User struct {
	UserId   string `json:"userId,omitempty" bson:"userId,omitempty"`
	Telefone string `json:"telefone,omitempty" bson:"-"`
	Name     string `json:"name,omitempty" bson:"-"`
	Email    string `json:"email,omitempty" bson:"-"`
	PassWord string `json:"password,omitempty" bson:"-"`
}
