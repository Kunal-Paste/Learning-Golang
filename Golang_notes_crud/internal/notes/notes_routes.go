package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoute(r *gin.Engine, db *mongo.Database) {
	repo := NewRepo(db)
	h := NewHandler(repo)

	noteGroup := r.Group("/notes")
	{
		noteGroup.POST("", h.CreateNote)
		noteGroup.GET("", h.ListNotes)
		noteGroup.GET("/:id", h.GetNoteByID)
		noteGroup.PUT("/:id", h.UpdateNoteByID)
		noteGroup.DELETE("/:id", h.DeleteNoteByID)
	}
}
