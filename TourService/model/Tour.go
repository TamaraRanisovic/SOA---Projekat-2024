package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Length      float64            `bson:"length,omitempty" json:"length"`
	Tags        []string           `bson:"tags,omitempty" json:"tags"`
	Difficulty  int                `bson:"difficulty,omitempty" json:"difficulty"`
	Price       float64            `bson:"price,omitempty" json:"price"`
	Guide_ID    string             `bson:"guide_id,omitempty" json:"guide_id"`
}

type Tours []*Tour

func (p *Tours) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tour) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tour) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}