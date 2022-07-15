package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"GRPC-project/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		Movies: &v1.Movies{
			MovieName:  "Doom",
			MovieGenre: "Action",
			Director:   "John Wachowski",
			Rating:     9.1,
		},
	}
	res1, err := c.CreateMovies(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	//Get all movies
	req2 := v1.ReadAllRequest{
		Api: apiVersion,
	}
	res2, err := c.GetAllMovies(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

	// Get movie by genre
	req3 := v1.ReadRequest{
		Api:        apiVersion,
		MovieGenre: "Action",
	}
	res3, err := c.GetMoviesByGenre(ctx, &req3)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res3)

	// Update
	req4 := v1.UpdateRequest{
		Api: apiVersion,
		Movies: &v1.Movies{
			MovieName:  "Doom",
			MovieGenre: "Action",
			Director:   "John Wachowski",
			Rating:     9.1,
		},
	}
	res4, err := c.UpdateMovies(ctx, &req4)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res4)

	// Delete
	req5 := v1.DeleteRequest{
		Api:       apiVersion,
		MovieName: "Doom",
	}
	res5, err := c.DeleteMovies(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}
