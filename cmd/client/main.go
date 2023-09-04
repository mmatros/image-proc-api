package main

import (
	"context"
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

	data, err := img.EncodeImage(im)
	if err != nil {
		log.Fatalf("can't encode image")
	}

	req := api.ConvertRequest{
		Image: data,
	}
	resp, err := client.ConvertImage(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to convert image %v", err)
	}
	data = resp.GetImage()
	result, err := img.ReadImageFromBuffer(data)
	if err != nil {
		log.Fatalf("failed to read image from buffer %v", err)
	}
	if err = img.SaveImage(result, "result.png"); err != nil {
		log.Fatalf("can't save image %v", err)
	}
}
