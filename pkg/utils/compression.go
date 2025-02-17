package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func Compress(input io.Reader, output io.Writer) error {
	// Utiliser un buffer pour stocker temporairement les données compressées
	var buffer bytes.Buffer

	// Créer un writer gzip qui écrit dans notre buffer
	gzipWriter := gzip.NewWriter(&buffer)

	// Copier les données d'entrée vers le gzip writer
	_, err := io.Copy(gzipWriter, input)
	if err != nil {
		return fmt.Errorf("erreur de copie: %v", err)
	}

	// Fermer le gzip writer pour s'assurer que toutes les données sont écrites
	if err := gzipWriter.Close(); err != nil {
		return fmt.Errorf("erreur de fermeture gzip: %v", err)
	}

	// Écrire le contenu du buffer vers la sortie
	_, err = io.Copy(output, &buffer)
	return err
}
