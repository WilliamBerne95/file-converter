package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/WilliamBerne95/file-converter/internal/converter"
	"github.com/WilliamBerne95/file-converter/pkg/utils"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()

	// Route de test
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Serveur de conversion actif")
	}).Methods("GET")

	// Routes de conversion
	r.HandleFunc("/convert/csv-to-json", handleCSVToJSON).Methods("POST")
	r.HandleFunc("/convert/image-to-jpeg", handleImageToJPEG).Methods("POST")

	// Route de compression (modification ici)
	r.HandleFunc("/compress", handleCompress).Methods("POST")

	return r
}

// Handler pour CSV vers JSON
func handleCSVToJSON(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	conv := &converter.TextConverter{}
	err = conv.CSVToJSON(file, w)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler pour la conversion d'image en JPEG
func handleImageToJPEG(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Début du traitement de la requête image-to-jpeg")

	// Limite la taille du fichier à 10MB
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du fichier: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Fichier reçu: %s, taille: %d bytes\n", header.Filename, header.Size)

	// Vérifier le type MIME
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du buffer", http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0) // Retour au début du fichier

	mimeType := http.DetectContentType(buffer)
	fmt.Printf("Type MIME détecté: %s\n", mimeType)

	if !strings.HasPrefix(mimeType, "image/") {
		http.Error(w, "Le fichier n'est pas une image", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	conv := &converter.ImageConverter{}
	if err := conv.ConvertToJPEG(file, w); err != nil {
		fmt.Printf("Erreur lors de la conversion: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de la conversion: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("Conversion terminée avec succès")
}

// Handler pour la compression
func handleCompress(w http.ResponseWriter, r *http.Request) {
	// Vérifier la méthode et le type de contenu
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Erreur de parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur de lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Configurer les headers de réponse
	w.Header().Set("Content-Type", "application/gzip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.gz", header.Filename))

	// Compresser le fichier
	if err := utils.Compress(file, w); err != nil {
		http.Error(w, "Erreur de compression: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
