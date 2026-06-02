package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (r *Repo) List(ctx context.Context) ([]Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{}

	cursor, err := r.coll.Find(opctx, filter)
	if err != nil {
		return nil, fmt.Errorf("find notes failed: %w", err)
	}

	defer cursor.Close(opctx)

	var note []Note

	if err := cursor.All(opctx, &note); err != nil {
		return nil, fmt.Errorf("decode notes failed: %v", err)
	}

	return note, nil

}
