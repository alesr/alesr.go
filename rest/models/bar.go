package models

import "gopkg.in/mgo.v2/bson"

type (

	// bar represents the structure of our resource
	Bar struct {
		ID          bson.ObjectId `json:"id" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		Description string        `json:"description" bson:"description"`
		Latitude    int           `json:"latitude" bson:"latitude"`
		Longitude   int           `json:"longitude" bson:"longitude"`
	}
)
