package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Lahoucine-7/gestion_articles_go/internal/controllers"
)

// SetupRoutes configure toutes les routes de l'application
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Groupe pour les articles
	articleRoutes := router.Group("/articles")
	{
		articleRoutes.POST("/", controllers.CreateArticle)
		articleRoutes.GET("/", controllers.GetArticles)
		articleRoutes.GET("/:id", controllers.GetArticle)
		articleRoutes.PUT("/:id", controllers.UpdateArticle)
		articleRoutes.DELETE("/:id", controllers.DeleteArticle)
	}

	return router
}
