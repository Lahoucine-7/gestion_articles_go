package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/Lahoucine-7/gestion_articles_go/internal/models"
)


var DB *gorm.DB

// InitDB initialise la connexion à la base de données
func InitDB() {
    // On pourrait récupérer les paramètres via des variables d'environnement
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Erreur de connexion à la base de données: %v", err)
    }
    DB = db
    // Auto-migration : Créer ou mettre à jour le schéma de la base de données
    if err := DB.AutoMigrate(&models.Article{}); err != nil {
        log.Fatalf("Erreur lors de la migration : %v", err)
    }
    log.Println("Connexion à la base de données établie et migration effectuée")
}
