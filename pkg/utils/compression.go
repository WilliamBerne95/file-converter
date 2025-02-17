package utils

import (
    "compress/gzip"
    "io"
)

func Compress(input io.Reader, output io.Writer) error {
    gzipWriter := gzip.NewWriter(output)
    defer gzipWriter.Close()

    _, err := io.Copy(gzipWriter, input)
    return err
}

func Decompress(input io.Reader, output io.Writer) error {
    gzipReader, err := gzip.NewReader(input)
    if err != nil {
        return err
    }
    defer gzipReader.Close()

    _, err = io.Copy(output, gzipReader)
    return err
}
