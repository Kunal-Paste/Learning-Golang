package server

import (
	"net/http"
	"note-api/internal/notes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Newrouter(database *mongo.Database) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "Healthy",
		})
	})

	notes.RegisterRoute(r, database)

	return r

}
