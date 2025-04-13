package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Lahoucine-7/gestion_articles_go/internal/models"
	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
)

// CreateArticle permet de créer un nouvel article
func CreateArticle(c *gin.Context) {
	var article models.Article

	// Bind des données JSON dans l'objet article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Création de l'article dans la base
	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'article"})
		return
	}
	c.JSON(http.StatusCreated, article)
}

// GetArticles permet de récupérer la liste des articles
func GetArticles(c *gin.Context) {
	var articles []models.Article
	if err := database.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des articles"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

// GetArticle permet de récupérer un article spécifique par son ID
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
		return
	}
	c.JSON(http.StatusOK, article)
}

// UpdateArticle permet de mettre à jour un article existant
func UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
		return
	}

	// Mise à jour des champs avec les nouvelles valeurs
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// DeleteArticle permet de supprimer un article
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
