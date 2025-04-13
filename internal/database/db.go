package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Lahoucine-7/gestion_articles_go/internal/models"
)

// DB est la variable globale qui contient la connexion à PostgreSQL.
var DB *gorm.DB

// InitDB initialise la connexion à la base de données en construisant le DSN à partir des variables d'environnement,
// puis effectue l'auto-migration du modèle Article.
func InitDB() {
	// Construction du DSN (Data Source Name) pour se connecter à PostgreSQL.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),    // Ex: localhost
		os.Getenv("DB_USER"),    // Ex: postgres (à améliorer pour utiliser un utilisateur dédié)
		os.Getenv("DB_PASSWORD"),// Mot de passe
		os.Getenv("DB_NAME"),    // Nom de la base, par ex. gestion_articles_test pour les tests
		os.Getenv("DB_PORT"),    // Ex: 5432
	)

	// Ouverture de la connexion avec GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}
	DB = db

	// Auto-migration du schéma pour créer ou mettre à jour la table "articles" à partir du modèle Article.
	if err := DB.AutoMigrate(&models.Article{}); err != nil {
		log.Fatalf("Erreur lors de la migration: %v", err)
	}

	log.Println("Connexion à la base de données établie et migration effectuée")
}
