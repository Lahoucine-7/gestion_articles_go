package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Lahoucine-7/gestion_articles_go/internal/controllers"
	swaggerFiles "github.com/swaggo/files"   // Fichiers Swagger embarqués
	ginSwagger "github.com/swaggo/gin-swagger" // Middleware pour Swagger
)

// SetupRoutes configure et retourne un routeur Gin intégrant l'ensemble des endpoints de l'application.
// Il sert également le dossier statique "docs" pour permettre l'accès aux fichiers Swagger générés.
func SetupRoutes() *gin.Engine {
	// Création d'un routeur Gin par défaut avec les middlewares Logger et Recovery.
	router := gin.Default()

	// Sert le dossier "docs" statiquement (assurez-vous d'exécuter l'application depuis la racine du projet).
	router.Static("/docs", "./docs")

	// Groupement des routes pour la gestion des articles.
	articleRoutes := router.Group("/articles")
	{
		articleRoutes.POST("/", controllers.CreateArticle)
		articleRoutes.GET("/", controllers.GetArticles)
		articleRoutes.GET("/:id", controllers.GetArticle)
		articleRoutes.PUT("/:id", controllers.UpdateArticle)
		articleRoutes.DELETE("/:id", controllers.DeleteArticle)
	}

	// Configuration de la route pour la documentation Swagger.
	// On utilise CustomWrapHandler pour spécifier l'URL absolue du fichier swagger.json.
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		URL: "http://localhost:8080/docs/swagger.json",
	}, swaggerFiles.Handler))

	return router
}
