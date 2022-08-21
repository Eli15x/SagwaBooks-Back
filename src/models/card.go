package models

type Card struct {
	UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
	Number string `json:"number,omitempty" bson:"number"`
	Name   string `json:"name,omitempty" bson:"-"`
	Data   string `json:"data,omitempty" bson:"-"`
}
