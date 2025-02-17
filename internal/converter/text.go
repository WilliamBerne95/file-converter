package converter

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

type TextConverter struct{}

// Structure générique pour les données
type Record map[string]string

// Structure pour XML
type XMLData struct {
	XMLName xml.Name `xml:"records"`
	Records []Record `xml:"record"`
}

// CSV vers JSON
func (c *TextConverter) CSVToJSON(csvFile io.Reader, jsonFile io.Writer) error {
	reader := csv.NewReader(csvFile)

	// Lire les en-têtes
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("erreur lecture en-têtes: %v", err)
	}

	var records []Record

	// Lire et convertir chaque ligne
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("erreur lecture ligne: %v", err)
		}

		record := make(Record)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}

	// Écrire en JSON avec indentation
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	return encoder.Encode(records)
}

// JSON vers CSV
func (c *TextConverter) JSONToCSV(jsonFile io.Reader, csvFile io.Writer) error {
	var records []Record

	// Lire le JSON
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&records); err != nil {
		return fmt.Errorf("erreur décodage JSON: %v", err)
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Extraire les en-têtes du premier enregistrement
	if len(records) == 0 {
		return fmt.Errorf("aucune donnée dans le JSON")
	}

	var headers []string
	for key := range records[0] {
		headers = append(headers, key)
	}

	// Écrire les en-têtes
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("erreur écriture en-têtes: %v", err)
	}

	// Écrire les données
	for _, record := range records {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = record[header]
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("erreur écriture ligne: %v", err)
		}
	}

	return nil
}

// CSV vers XML
func (c *TextConverter) CSVToXML(csvFile io.Reader, xmlFile io.Writer) error {
	reader := csv.NewReader(csvFile)

	// Lire les en-têtes
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("erreur lecture en-têtes: %v", err)
	}

	var records []Record

	// Lire et convertir chaque ligne
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("erreur lecture ligne: %v", err)
		}

		record := make(Record)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}

	// Créer la structure XML
	data := XMLData{Records: records}

	// Écrire l'en-tête XML
	xmlFile.Write([]byte(xml.Header))

	// Encoder en XML avec indentation
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "    ")
	return encoder.Encode(data)
}

// XML vers CSV
func (c *TextConverter) XMLToCSV(xmlFile io.Reader, csvFile io.Writer) error {
	var data XMLData

	// Décoder le XML
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("erreur décodage XML: %v", err)
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	if len(data.Records) == 0 {
		return fmt.Errorf("aucune donnée dans le XML")
	}

	// Extraire les en-têtes
	var headers []string
	for key := range data.Records[0] {
		headers = append(headers, key)
	}

	// Écrire les en-têtes
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("erreur écriture en-têtes: %v", err)
	}

	// Écrire les données
	for _, record := range data.Records {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = record[header]
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("erreur écriture ligne: %v", err)
		}
	}

	return nil
}

// JSON vers XML
func (c *TextConverter) JSONToXML(jsonFile io.Reader, xmlFile io.Writer) error {
	var records []Record

	// Lire le JSON
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&records); err != nil {
		return fmt.Errorf("erreur décodage JSON: %v", err)
	}

	// Créer la structure XML
	data := XMLData{Records: records}

	// Écrire l'en-tête XML
	xmlFile.Write([]byte(xml.Header))

	// Encoder en XML avec indentation
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "    ")
	return encoder.Encode(data)
}

// XML vers JSON
func (c *TextConverter) XMLToJSON(xmlFile io.Reader, jsonFile io.Writer) error {
	var data XMLData

	// Décoder le XML
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("erreur décodage XML: %v", err)
	}

	// Encoder en JSON avec indentation
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data.Records)
}
