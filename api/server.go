package api

type Server struct {
	UnimplementedImageProcApiServer
}

func NewServer() Server {
	return Server{}
}
