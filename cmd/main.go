package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
	"github.com/Lahoucine-7/gestion_articles_go/internal/routes"
	_ "github.com/Lahoucine-7/gestion_articles_go/docs" // Import anonyme pour initialiser Swagger via swaggo
)

// main est le point d'entrée de l'application
func main() {
	// Charger les variables d'environnement depuis un fichier .env, le cas échéant.
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement existantes")
	}

	// Initialisation de la connexion à la base de données et exécution de l'auto-migration.
	database.InitDB()

	// Configuration des routes de l'API.
	router := routes.SetupRoutes()

	// Démarrage du serveur sur le port 8080.
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
