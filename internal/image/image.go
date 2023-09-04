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

// EncodeImage ...
func EncodeImage(image image.Image) ([]byte, error) {
	buff := new(bytes.Buffer)

	err := png.Encode(buff, image)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// SaveImage ...
func SaveImage(im image.Image, filePath string) error {
	data, err := EncodeImage(im)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}
