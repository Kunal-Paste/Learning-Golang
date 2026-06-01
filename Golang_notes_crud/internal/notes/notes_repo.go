package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {
	opCxt, cancle := context.WithTimeout(ctx, 5*time.Second)
	defer cancle()

	_, err := r.coll.InsertOne(opCxt, note)
	if err != nil {
		return Note{}, fmt.Errorf("insert note failed")
	}

	return note, nil
}
