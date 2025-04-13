package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
	"github.com/Lahoucine-7/gestion_articles_go/internal/models"
)

// CreateArticle gère la création d'un nouvel article dans la base de données.
//
// @Summary Créer un nouvel article
// @Description Crée un nouvel article en enregistrant le titre, le contenu, l'auteur et la date de publication.
// @Tags Articles
// @Accept json
// @Produce json
// @Param article body models.Article true "Les données de l'article à créer"
// @Success 201 {object} models.Article "Article créé avec succès"
// @Failure 400 {object} map[string]string "Erreur de validation des données"
// @Failure 500 {object} map[string]string "Erreur lors de la création de l'article"
// @Router /articles [post]
func CreateArticle(c *gin.Context) {
	var article models.Article

	// Liaison automatique du JSON reçu dans la requête avec la structure Article.
	if err := c.ShouldBindJSON(&article); err != nil {
		// Réponse en cas d'erreur de liaison ou de validation.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Création de l'article dans la base de données via GORM.
	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'article"})
		return
	}

	// Réponse avec le statut 201 Created et l'article créé.
	c.JSON(http.StatusCreated, article)
}

// GetArticles récupère la liste complète des articles.
//
// @Summary Récupérer la liste des articles
// @Description Récupère tous les articles enregistrés dans la base de données.
// @Tags Articles
// @Produce json
// @Success 200 {array} models.Article "Liste des articles"
// @Failure 500 {object} map[string]string "Erreur lors de la récupération"
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	var articles []models.Article

	// Exécution d'une requête pour récupérer tous les articles.
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des articles"})
		return
	}

	// Réponse avec la liste des articles.
	c.JSON(http.StatusOK, articles)
}

// GetArticle récupère un article spécifique à partir de son identifiant.
//
// @Summary Récupérer un article par son ID
// @Description Récupère les détails d'un article en se basant sur l'ID fourni dans l'URL.
// @Tags Articles
// @Produce json
// @Param id path int true "ID de l'article"
// @Success 200 {object} models.Article "Détails de l'article"
// @Failure 400 {object} map[string]string "ID invalide"
// @Failure 404 {object} map[string]string "Article non trouvé"
// @Router /articles/{id} [get]
func GetArticle(c *gin.Context) {
	// Conversion de l'ID reçu dans l'URL en entier.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var article models.Article
	// Recherche d'un article ayant l'ID spécifié.
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
		return
	}

	// Réponse avec l'article trouvé.
	c.JSON(http.StatusOK, article)
}

// UpdateArticle met à jour les informations d'un article existant.
//
// @Summary Mettre à jour un article existant
// @Description Met à jour un article identifié par son ID en utilisant les données fournies dans le corps de la requête.
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "ID de l'article"
// @Param article body models.Article true "Données mises à jour pour l'article"
// @Success 200 {object} models.Article "Article mis à jour"
// @Failure 400 {object} map[string]string "ID invalide ou erreur de validation"
// @Failure 404 {object} map[string]string "Article non trouvé"
// @Failure 500 {object} map[string]string "Erreur lors de la mise à jour"
// @Router /articles/{id} [put]
func UpdateArticle(c *gin.Context) {
	// Conversion de l'ID provenant des paramètres de l'URL.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var article models.Article
	// Récupération préalable de l'article afin de vérifier son existence.
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
		return
	}

	// Lier les données JSON reçues afin de mettre à jour l'article.
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sauvegarde des modifications dans la base de données.
	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'article"})
		return
	}

	// Réponse avec l'article mis à jour.
	c.JSON(http.StatusOK, article)
}

// DeleteArticle supprime un article de la base de données.
//
// @Summary Supprimer un article
// @Description Supprime un article identifié par son ID de la base de données.
// @Tags Articles
// @Produce json
// @Param id path int true "ID de l'article"
// @Success 204 {object} map[string]string "Article supprimé"
// @Failure 400 {object} map[string]string "ID invalide"
// @Failure 500 {object} map[string]string "Erreur lors de la suppression"
// @Router /articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	// Conversion et vérification de l'ID fourni.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	// Suppression de l'article dans la base de données.
	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'article"})
		return
	}

	// Réponse avec le statut 204 No Content pour indiquer le succès de la suppression.
	c.JSON(http.StatusNoContent, gin.H{})
}
