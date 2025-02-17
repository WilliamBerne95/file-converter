package main

import (
    "log"
    "net/http"
    "github.com/yourusername/file-converter/internal/api"
    "github.com/yourusername/file-converter/internal/cli"
)

func main() {
    // Démarrer le serveur API
    go func() {
        log.Fatal(http.ListenAndServe(":8080", api.Router()))
    }()

    // Exécuter l'interface CLI
    cli.Execute()
}