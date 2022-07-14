package apis

import (
	"github.com/jmoiron/sqlx"
)

const (
	apiVersion = "v1"
)

type MovieServiceServer struct {
	db *sqlx.DB
}

func NewMovieServiceServer(db *sqlx.DB) *MovieServiceServer {
	return &MovieServiceServer{db: db}
}

func (s *MovieServiceServer) GetVersion() string {
	return apiVersion
}
