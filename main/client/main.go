package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"GRPC-project/pkg/api/proto/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "movie"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:9090", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewMoviesServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	//// Call Create
	//req1 := v1.CreateRequest{
	//	Api: apiVersion,
	//	Movies: &v1.Movies{
	//		MovieName:  config.MoviesName[rand.Intn(len(config.MoviesName))],
	//		MovieGenre: config.MoviesGenre[rand.Intn(len(config.MoviesGenre))],
	//		Director:   config.MoviesDirector[rand.Intn(len(config.MoviesDirector))],
	//		Rating:     float32(config.MoviesRating[rand.Intn(len(config.MoviesRating))]),
	//	},
	//}
	//res1, err := c.CreateMovies(ctx, &req1)
	//if err != nil {
	//	log.Fatalf("Create failed: %v", err)
	//}
	//log.Printf("Create result: <%+v>\n\n", res1)

	//Get movies by genre
	req2 := v1.ReadRequest{
		Api:        apiVersion,
		MovieGenre: "Action",
	}
	res2, err := c.GetMovieByGenre(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	fmt.Println("Get movies by genre:")
	for _, m := range res2.Movies {
		log.Printf("Movie: <%+v>\n", m)
	}
	//
	////Get all movies
	//req3 := v1.ReadAllRequest{
	//	Api: apiVersion,
	//}
	//res3, err := c.GetAllMovies(ctx, &req3)
	//if err != nil {
	//	log.Fatalf("Read failed: %v", err)
	//}
	////log.Printf("Read result: <%+v>\n\n", res3)
	//fmt.Println("Get all movies:")
	//for _, m := range res3.Movies {
	//	log.Printf("Movie: <%+v>\n", m)
	//}
	//
	//// Update
	//req4 := v1.UpdateRequest{
	//	Api:       apiVersion,
	//	MovieName: config.MoviesName[rand.Intn(len(config.MoviesName))],
	//}
	//res4, err := c.UpdateMovies(ctx, &req4)
	//if err != nil {
	//	log.Fatalf("Update failed: %v", err)
	//}
	//log.Printf("Update successfull: <%+v>\n\n", res4)
	//
	//// Delete
	//req5 := v1.DeleteRequest{
	//	Api:       apiVersion,
	//	MovieName: "Doom",
	//}
	//res5, err := c.DeleteMovies(ctx, &req5)
	//if err != nil {
	//	log.Fatalf("Delete failed: %v", err)
	//}
	//log.Printf("Delete successfull: <%+v>\n\n", res5)
}
