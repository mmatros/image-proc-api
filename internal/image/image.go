package image

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"os"
)

// ReadImageFromFile ...
func ReadImageFromFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readImage(f)
}

// ReadImageFromBuffer ...
func ReadImageFromBuffer(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	return readImage(reader)
}

func readImage(reader io.Reader) (image.Image, error) {
	image, _, err := image.Decode(reader)
	return image, err
}

// SaveImage ...
func SaveImage(im image.Image, filePath string) error {
	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, im)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	_, err = file.Write(buff.Bytes())
	return err
}
