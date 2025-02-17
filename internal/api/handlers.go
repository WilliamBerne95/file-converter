package api

import (
	"fmt"
	"net/http"
	"path/filepath"
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
	r.HandleFunc("/convert/json-to-csv", handleJSONToCSV).Methods("POST")
	r.HandleFunc("/convert/csv-to-xml", handleCSVToXML).Methods("POST")
	r.HandleFunc("/convert/xml-to-csv", handleXMLToCSV).Methods("POST")
	r.HandleFunc("/convert/json-to-xml", handleJSONToXML).Methods("POST")
	r.HandleFunc("/convert/xml-to-json", handleXMLToJSON).Methods("POST")

	// Route de compression (modification ici)
	r.HandleFunc("/compress", handleCompress).Methods("POST")
	r.HandleFunc("/decompress", handleDecompress).Methods("POST")

	return r
}

// Handler pour la conversion d'image en JPEG
func handleImageToJPEG(w http.ResponseWriter, r *http.Request) {
	// Vérifier la taille maximale
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Fichier trop volumineux", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Conversion de l'image: %s\n", fileHeader.Filename)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jpg",
		strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))))

	converter := &converter.ImageConverter{}
	if err := converter.ConvertToJPEG(file, w); err != nil {
		http.Error(w, fmt.Sprintf("Erreur de conversion: %v", err), http.StatusInternalServerError)
		return
	}
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

func handleDecompress(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Retirer l'extension .gz du nom du fichier original
	originalName := strings.TrimSuffix(header.Filename, ".gz")
	w.Header().Set("Content-Disposition", "attachment; filename="+originalName)
	w.Header().Set("Content-Type", "application/octet-stream")

	if err := utils.Decompress(file, w); err != nil {
		http.Error(w, "Erreur lors de la décompression: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleCSVToJSON(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	conv := &converter.TextConverter{}
	if err := conv.CSVToJSON(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleJSONToCSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/csv")
	conv := &converter.TextConverter{}
	if err := conv.JSONToCSV(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleCSVToXML(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/xml")
	conv := &converter.TextConverter{}
	if err := conv.CSVToXML(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleXMLToCSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/csv")
	conv := &converter.TextConverter{}
	if err := conv.XMLToCSV(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleJSONToXML(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/xml")
	conv := &converter.TextConverter{}
	if err := conv.JSONToXML(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleXMLToJSON(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erreur lecture fichier: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	conv := &converter.TextConverter{}
	if err := conv.XMLToJSON(file, w); err != nil {
		http.Error(w, "Erreur conversion: "+err.Error(), http.StatusInternalServerError)
	}
}
