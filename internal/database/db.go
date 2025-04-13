package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Lahoucine-7/gestion_articles_go/internal/models"
)

// DB est la variable globale qui stocke la connexion à la base de données.
var DB *gorm.DB

// InitDB initialise la connexion à PostgreSQL et lance l'auto-migration du modèle Article.
// Il construit le Data Source Name (DSN) à partir des variables d'environnement.
func InitDB() {
	// Construction du DSN pour PostgreSQL.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),    // ex: localhost
		os.Getenv("DB_USER"),    // ex: articles_admin
		os.Getenv("DB_PASSWORD"),// ex: ton mot de passe
		os.Getenv("DB_NAME"),    // ex: gestion_articles
		os.Getenv("DB_PORT"),    // ex: 5432
	)

	// Ouverture de la connexion avec GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}
	DB = db

	// Exécution de l'auto-migration pour créer ou mettre à jour la table articles.
	if err := DB.AutoMigrate(&models.Article{}); err != nil {
		log.Fatalf("Erreur lors de la migration: %v", err)
	}

	log.Println("Connexion à la base de données établie et migration effectuée")
}
