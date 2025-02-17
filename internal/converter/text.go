package converter

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

type TextConverter struct{}

func (c *TextConverter) CSVToJSON(csvFile io.Reader, jsonFile io.Writer) error {
	reader := csv.NewReader(csvFile)

	// Lire les en-têtes
	headers, err := reader.Read()
	if err != nil {
		log.Printf("Erreur lors de la lecture des en-têtes: %v", err)
		return err
	}
	log.Printf("En-têtes lus: %v", headers)

	var records []map[string]string

	// Lire et convertir chaque ligne
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Erreur lors de la lecture d'une ligne: %v", err)
			return err
		}
		log.Printf("Ligne lue: %v", row)

		record := make(map[string]string)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}

	log.Printf("Nombre total d'enregistrements: %d", len(records))

	// Écrire en JSON avec une indentation pour meilleure lisibilité
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(records)
	if err != nil {
		log.Printf("Erreur lors de l'encodage JSON: %v", err)
		return err
	}
	log.Printf("Conversion terminée avec succès")
	return nil
}
