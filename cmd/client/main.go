package main

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"log"

	img "github.com/mmatros/image-proc-api/internal/image"

	api "github.com/mmatros/image-proc-api/pkg/api/imageproc_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	im, err := img.ReadImageFromFile("lena.png")
	if err != nil {
		log.Fatalf("error on read image from file")
	}

	conn, err := grpc.Dial(
		"localhost:12000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewImageProcApiClient(conn)

	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err = png.Encode(buff, im)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}
	req := api.ConvertRequest{
		Image: buff.Bytes(),
	}
	resp, err := client.ConvertImage(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to convert image %v", err)
	}
	data := resp.GetImage()
	result, err := img.ReadImageFromBuffer(data)
	if err != nil {
		log.Fatalf("failed to read image from buffer %v", err)
	}
	img.SaveImage(result, "result.png")
}
