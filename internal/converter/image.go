package converter

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"io"
	"net/http" // Ajout pour DetectContentType
)

type ImageConverter struct{}

func (c *ImageConverter) ConvertToJPEG(input io.Reader, output io.Writer) error {
	// Lire l'image en mémoire
	imageBytes, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("erreur de lecture: %v", err)
	}

	// Détecter le type de contenu
	contentType := http.DetectContentType(imageBytes)
	fmt.Printf("Type de contenu détecté: %s\n", contentType)

	// Vérifier si c'est une image
	if !bytes.HasPrefix(imageBytes[:8], []byte("\x89PNG")) &&
		!bytes.HasPrefix(imageBytes[:3], []byte("GIF")) &&
		!bytes.HasPrefix(imageBytes[:2], []byte("\xFF\xD8")) {
		return fmt.Errorf("format de fichier non supporté: %s", contentType)
	}

	// Décoder l'image
	img, format, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return fmt.Errorf("erreur de décodage format %s: %v", format, err)
	}

	fmt.Printf("Format d'image détecté: %s\n", format)

	// Encoder en JPEG
	return jpeg.Encode(output, img, &jpeg.Options{
		Quality: 90,
	})
}
