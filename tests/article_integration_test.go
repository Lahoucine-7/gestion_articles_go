// test/article_integration_test.go
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
	"github.com/Lahoucine-7/gestion_articles_go/internal/models"
	"github.com/Lahoucine-7/gestion_articles_go/internal/routes"
)

// setupTestRouter initialise l'environnement de test en configurant les variables d'environnement par défaut,
// en initialisant la base de données et en nettoyant la table articles.
func setupTestRouter() *gin.Engine {
	// Configurer des valeurs par défaut pour les tests.
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "admin")
	// Pour les tests, il est judicieux d'utiliser une base spécifique (gestion_articles_test par exemple)
	os.Setenv("DB_NAME", "gestion_articles_test")
	os.Setenv("DB_PORT", "5432")

	// Initialiser la base de données.
	database.InitDB()

	// Nettoyer la table articles pour un environnement de test isolé.
	database.DB.Exec("DELETE FROM articles")

	// Retourner le routeur configuré.
	return routes.SetupRoutes()
}

// TestCreateArticle vérifie la création d'un nouvel article via l'endpoint POST /articles.
func TestCreateArticle(t *testing.T) {
	router := setupTestRouter()

	article := models.Article{
		Title:         "Test Article",
		Content:       "This is the test content.",
		Author:        "Test Author",
		PublishedDate: time.Now(),
	}
	jsonData, err := json.Marshal(article)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/articles/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdArticle models.Article
	err = json.Unmarshal(w.Body.Bytes(), &createdArticle)
	assert.NoError(t, err)
	assert.NotZero(t, createdArticle.ID)
	assert.Equal(t, article.Title, createdArticle.Title)
}

// TestGetArticles vérifie que l'endpoint GET /articles retourne au moins un article.
func TestGetArticles(t *testing.T) {
	router := setupTestRouter()

	testArticle := models.Article{
		Title:         "Integration Test Article",
		Content:       "Content for integration test.",
		Author:        "Integration Author",
		PublishedDate: time.Now(),
	}
	err := database.DB.Create(&testArticle).Error
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/articles/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var articles []models.Article
	err = json.Unmarshal(w.Body.Bytes(), &articles)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(articles), 1)
}

// TestGetArticle vérifie la récupération d'un article précis via GET /articles/{id}.
func TestGetArticle(t *testing.T) {
	router := setupTestRouter()

	testArticle := models.Article{
		Title:         "Unique Article",
		Content:       "Unique Content",
		Author:        "Unique Author",
		PublishedDate: time.Now(),
	}
	err := database.DB.Create(&testArticle).Error
	assert.NoError(t, err)

	url := "/articles/" + strconv.Itoa(int(testArticle.ID))
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedArticle models.Article
	err = json.Unmarshal(w.Body.Bytes(), &fetchedArticle)
	assert.NoError(t, err)
	assert.Equal(t, testArticle.Title, fetchedArticle.Title)
}

// TestUpdateArticle vérifie la mise à jour d'un article via PUT /articles/{id}.
func TestUpdateArticle(t *testing.T) {
	router := setupTestRouter()

	testArticle := models.Article{
		Title:         "Old Title",
		Content:       "Old Content",
		Author:        "Old Author",
		PublishedDate: time.Now(),
	}
	err := database.DB.Create(&testArticle).Error
	assert.NoError(t, err)

	testArticle.Title = "New Title"
	updatedJSON, err := json.Marshal(testArticle)
	assert.NoError(t, err)

	url := "/articles/" + strconv.Itoa(int(testArticle.ID))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(updatedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var updatedArticle models.Article
	err = json.Unmarshal(w.Body.Bytes(), &updatedArticle)
	assert.NoError(t, err)
	assert.Equal(t, "New Title", updatedArticle.Title)
}

// TestDeleteArticle vérifie que l'endpoint DELETE /articles/{id} supprime bien l'article.
func TestDeleteArticle(t *testing.T) {
	router := setupTestRouter()

	testArticle := models.Article{
		Title:         "Delete Me",
		Content:       "Content to Delete",
		Author:        "Delete Author",
		PublishedDate: time.Now(),
	}
	err := database.DB.Create(&testArticle).Error
	assert.NoError(t, err)

	url := "/articles/" + strconv.Itoa(int(testArticle.ID))
	req, _ := http.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var fetchedArticle models.Article
	err = database.DB.First(&fetchedArticle, testArticle.ID).Error
	assert.Error(t, err) // L'article ne doit pas être trouvé
}
