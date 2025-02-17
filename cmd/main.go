package main

import (
	"log"
	"net/http"

	"github.com/WilliamBerne95/file-converter/internal/api"
)

func main() {
	// Log de démarrage
	log.Printf("Démarrage du serveur...")

	// Création du router
	router := api.Router()

	// Démarrage du serveur
	log.Printf("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
