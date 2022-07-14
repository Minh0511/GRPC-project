package movie

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Receive message from client: %s", in.Name)
	return &HelloReply{Message: "Hello From the Server!"}, nil
}
