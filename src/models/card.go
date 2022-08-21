package models

type Card struct {
	UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
	Numero string `json:"numero,omitempty" bson:"numero"`
	Name   string `json:"name,omitempty" bson:"-"`
	Data   string `json:"data,omitempty" bson:"-"`
}
