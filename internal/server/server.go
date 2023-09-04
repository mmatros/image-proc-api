package server

import (
	api "github.com/mmatros/image-proc-api/pkg/api/imageproc_v1"
)

type Server struct {
	api.UnimplementedImageProcApiServer
}

func NewServer() Server {
	return Server{}
}
