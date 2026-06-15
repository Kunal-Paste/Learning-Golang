package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userService := user.NewService(userRepo, a.Config.JWTsecret)

	userHandler := user.NewHandler(userService)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	api := r.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTsecret))

	api.GET("/files", func(c *gin.Context) {

		userID, _ := middleware.GetUserID(c)

		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"userId": userID,
			"file":   []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"file": []any{},
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func(c *gin.Context) {

		role, _ := middleware.GetRole(c)

		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"role": role,
		})
	})

	return r
}
