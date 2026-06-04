package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context) {

	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	now := time.Now().UTC()

	note := Note{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.Create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create note here !",
		})

		return
	}

	c.JSON(http.StatusCreated, created)

}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch all notes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func (h *Handler) GetNoteByID(c *gin.Context) {

	idStr := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json id",
		})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), objId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "note not found for the given id",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch the note",
		})

		return
	}

	c.JSON(http.StatusOK, note)

}

func (h *Handler) UpdateNoteByID(c *gin.Context) {
	idstr := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json id",
		})
		return
	}

	var req UpdateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid josn format",
		})
		return
	}

	updated, err := h.repo.UpdateByID(c.Request.Context(), objId, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "note not found for the given id",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch the note",
		})

		return
	}

	c.JSON(http.StatusOK, updated)

}

func (h *Handler) DeleteNoteByID(c *gin.Context) {
	idstr := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json id",
		})
		return
	}

	deleted, err := h.repo.DeleteByID(c.Request.Context(), objId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete the node",
		})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

}
