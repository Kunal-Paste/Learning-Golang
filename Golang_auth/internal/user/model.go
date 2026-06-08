package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Email string `bson:"email" json:"email"`

	PasswordHash string `bson:"passwordhash" json:"-"`

	Role string `bson:"role" json:"role"`

	CreatedAt time.Time `bson:"CreatedAt" json:"CreatedAt"`

	UpdatedAt time.Time `bson:"UpdatedAt" json:"UpdatedAt"`
}

type PublicUser struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Role string `json:"role"`

	CreatedAt time.Time `json:"CreatedAt"`

	UpdatedAt time.Time `json:"UpdatedAt"`
}

func ToPublic(u User) PublicUser {
	return PublicUser{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
