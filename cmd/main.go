package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
	"github.com/Lahoucine-7/gestion_articles_go/internal/routes"
	_ "github.com/Lahoucine-7/gestion_articles_go/docs" // Import anonyme pour initialiser la documentation Swagger générée par swag
)

// main est le point d'entrée de l'application.
// Il charge les variables d'environnement, initialise la DB, et démarre le serveur HTTP.
func main() {
	// Charger le fichier .env (s'il existe) pour définir les variables d'environnement.
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement existantes")
	}

	// Initialisation de la connexion à la base de données et auto-migration du modèle Article.
	database.InitDB()

	// Configuration et récupération du routeur de l'API.
	router := routes.SetupRoutes()

	// Démarrage du serveur HTTP sur le port 8080.
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
