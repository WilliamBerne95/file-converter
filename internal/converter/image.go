package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

type ImageConverter struct{}

func (c *ImageConverter) ConvertToJPEG(input io.Reader, output io.Writer) error {
	// Lire le contenu du fichier
	imgData, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("erreur de lecture: %v", err)
	}

	// Détecter le type de contenu
	contentType := http.DetectContentType(imgData)
	fmt.Printf("Type de contenu détecté: %s\n", contentType)

	// Décoder l'image selon son type
	var img image.Image
	inputReader := bytes.NewReader(imgData)

	switch contentType {
	case "image/png":
		img, err = png.Decode(inputReader)
	case "image/jpeg":
		img, err = jpeg.Decode(inputReader)
	default:
		return fmt.Errorf("format d'image non supporté: %s", contentType)
	}

	if err != nil {
		return fmt.Errorf("erreur de décodage: %v", err)
	}

	// Encoder en JPEG
	return jpeg.Encode(output, img, &jpeg.Options{Quality: 85})
}
