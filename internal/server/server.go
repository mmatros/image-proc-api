package server

import (
	"context"

	img "github.com/mmatros/image-proc-api/internal/image"
	api "github.com/mmatros/image-proc-api/pkg/api/imageproc_v1"
)

type Server struct {
	api.UnimplementedImageProcApiServer
}

func NewServer() *Server {
	return &Server{}
}

// ConvertImage ...
func (s *Server) ConvertImage(ctx context.Context, req *api.ConvertRequest) (*api.ConvertResponse, error) {
	inputData := req.GetImage()
	image, err := img.ReadImageFromBuffer(inputData)
	if err != nil {
		return nil, err
	}
	var resp api.ConvertResponse
	resp.Image, err = img.EncodeImage(image)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
