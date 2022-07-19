package v1

import (
	"GRPC-project/pkg/api/proto/v1"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "movie"
)

type movieServiceServer struct {
	v1.UnimplementedMoviesServiceServer
	db *sqlx.DB
}

func NewMovieServiceServer(db *sqlx.DB) v1.MoviesServiceServer {
	return &movieServiceServer{db: db}
}

func (s *movieServiceServer) CreateMovies(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO Movies (MovieName, MovieGenre, Director, Rating) VALUES (?, ?, ?, ?)", request.Movies.MovieName, request.Movies.MovieGenre, request.Movies.Director, request.Movies.Rating)
	tx.Commit()
	return &v1.CreateResponse{
		Api:    apiVersion,
		Movies: request.Movies,
	}, nil
}

func (s *movieServiceServer) GetAllMovies(ctx context.Context, request *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	rows, err := s.db.Queryx("SELECT MovieName, MovieGenre, Director, Rating FROM Movies")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to query-> "+err.Error())
	}

	var movies []*v1.Movies
	for rows.Next() {
		var m v1.Movies
		if err := rows.Scan(&m.MovieName, &m.MovieGenre, &m.Director, &m.Rating); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from row-> "+err.Error())
		}
		movies = append(movies, &m)
	}
	return &v1.ReadAllResponse{
		Api:    apiVersion,
		Movies: movies,
	}, nil
}

type responseForGetMovieByName struct {
	ID         int     `db:"ID"`
	MovieName  string  `db:"MovieName"`
	MovieGenre string  `db:"MovieGenre"`
	Director   string  `db:"Director"`
	Rating     float32 `db:"Rating"`
}

func (s *movieServiceServer) GetMovieByGenre(ctx context.Context, request *v1.ReadRequest) (*v1.ReadResponse, error) {
	start := time.Now()
	tx := s.db.MustBegin()
	listMovie := "SELECT ID, MovieName, MovieGenre, Director, Rating FROM Movies WHERE MovieGenre = ?"
	var queryAns []*responseForGetMovieByName
	err := tx.SelectContext(ctx, &queryAns, listMovie, request.MovieGenre)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to query-> "+err.Error())
	}

	var movies []*v1.Movies
	for _, v := range queryAns {
		var m v1.Movies
		m.ID = int32(v.ID)
		m.MovieName = v.MovieName
		m.MovieGenre = v.MovieGenre
		m.Director = v.Director
		m.Rating = v.Rating
		movies = append(movies, &m)
	}
	tx.Commit()
	fmt.Println("Get movie by genre executed in: ", time.Since(start))
	return &v1.ReadResponse{
		Api:    apiVersion,
		Movies: movies,
	}, nil
}

func (s *movieServiceServer) UpdateMovies(ctx context.Context, request *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("UPDATE Movies SET MovieGenre = ?, Director = ?, Rating = ? WHERE MovieName = ?", "Action", "Le Thi Linh Chi", 20, request.MovieName)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateResponse{
		Api: apiVersion,
	}, nil
}

func (s *movieServiceServer) DeleteMovies(ctx context.Context, request *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("DELETE FROM Movies WHERE MovieName = ?", request.MovieName)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}
	return &v1.DeleteResponse{
		Api: apiVersion,
	}, nil
}

// checkAPI checks if the API version requested by client is supported by server
func (s *movieServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *movieServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
