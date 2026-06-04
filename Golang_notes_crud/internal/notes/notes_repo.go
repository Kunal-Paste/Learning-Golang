package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var note Note

	err := r.coll.FindOne(opctx, filter, options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("find by id failed : %w", err)
	}

	return note, nil
}

func (r *Repo) UpdateByID(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title":     req.Title,
			"content":   req.Content,
			"pinned":    req.Pinned,
			"UpdatedAt": time.Now().UTC(),
		},
	}

	after := options.After

	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var updated Note

	err := r.coll.FindOneAndUpdate(opctx, filter, update, &opts).Decode(&updated)
	if err != nil {
		return Note{}, fmt.Errorf("update note failed : %w", err)
	}

	return updated, nil

}

func (r *Repo) DeleteByID(ctx context.Context, id primitive.ObjectID) (bool, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	res, err := r.coll.DeleteOne(opctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to delete the given note : %w", err)
	}

	if res.DeletedCount == 0 {
		return false, nil
	}

	return true, nil

}
