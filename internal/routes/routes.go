package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Lahoucine-7/gestion_articles_go/internal/controllers"
	swaggerFiles "github.com/swaggo/files"     // Fichiers Swagger embarqués
	ginSwagger "github.com/swaggo/gin-swagger"   // Middleware pour Swagger
)

// SetupRoutes configure et retourne un routeur Gin avec toutes les routes de l'application.
func SetupRoutes() *gin.Engine {
	// Création d'un nouveau routeur par défaut avec Logger et Recovery intégrés.
	router := gin.Default()

	// Sert le dossier "docs" de façon statique pour rendre accessible swagger.json et les autres fichiers.
	// Le chemin "docs" est relatif à la racine d'exécution du projet.
	router.Static("/docs", "./docs")

	// Configuration des routes pour la gestion des articles.
	articleRoutes := router.Group("/articles")
	{
		articleRoutes.POST("/", controllers.CreateArticle)
		articleRoutes.GET("/", controllers.GetArticles)
		articleRoutes.GET("/:id", controllers.GetArticle)
		articleRoutes.PUT("/:id", controllers.UpdateArticle)
		articleRoutes.DELETE("/:id", controllers.DeleteArticle)
	}

	// Configuration de la route pour Swagger.
	// On utilise CustomWrapHandler avec une URL absolue pour pointer vers swagger.json dans le dossier "docs".
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		URL: "http://localhost:8080/docs/swagger.json",
	}, swaggerFiles.Handler))

	return router
}
