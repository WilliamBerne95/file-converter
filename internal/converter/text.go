package converter

import (
    "encoding/csv"
    "encoding/json"
    "encoding/xml"
    "io"
    "os"
)

type TextConverter struct{}

func (c *TextConverter) CSVToJSON(csvFile io.Reader, jsonFile io.Writer) error {
    reader := csv.NewReader(csvFile)
    
    // Lire les en-têtes
    headers, err := reader.Read()
    if err != nil {
        return err
    }

    var records []map[string]string
    
    // Lire et convertir chaque ligne
    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        record := make(map[string]string)
        for i, value := range row {
            record[headers[i]] = value
        }
        records = append(records, record)
    }

    // Écrire en JSON
    encoder := json.NewEncoder(jsonFile)
    return encoder.Encode(records)
}