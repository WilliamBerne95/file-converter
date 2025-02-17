package converter

import (
    "image"
    "image/jpeg"
    "image/png"
    "image/gif"
    "io"
    "os"
)

type ImageConverter struct{}

func (c *ImageConverter) ConvertToJPEG(input io.Reader, output io.Writer) error {
    img, _, err := image.Decode(input)
    if err != nil {
        return err
    }
    return jpeg.Encode(output, img, nil)
}