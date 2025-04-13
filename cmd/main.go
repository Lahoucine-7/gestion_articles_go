package main

import (
	"log"
	"github.com/Lahoucine-7/gestion_articles_go/internal/database"
	"github.com/Lahoucine-7/gestion_articles_go/internal/routes"
)

func main() {
	// Initialiser la base de données et effectuer l'auto-migration
	database.InitDB()

	// Configurer le routeur (et par extension, toutes les routes)
	router := routes.SetupRoutes()

	// Démarrer le serveur sur le port 8080 (modifiable selon tes besoins)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}