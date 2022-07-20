package main

import (
	v1 "GRPC-project/pkg/api/proto/v1"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "movie"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	redisClient := newRedisClient()

	searchMovieGenre := "Action"

	start := time.Now()
	nameInRedis, err := redisClient.HGetAll(ctx, searchMovieGenre).Result()
	end := time.Since(start)

	if err != nil {
		log.Println(err)
		panic(err)
	}
	res := addDBtoRedis(ctx, redisClient, searchMovieGenre)
	if res != nil {
		log.Println(res)
		panic(res)
	}
	if len(nameInRedis) == 0 {
		log.Println("No movie in redis cache found")

	} else {
		log.Println("Found movie in redis")
		for key, value := range nameInRedis {
			log.Println("ID:", key, value)
		}
	}
	fmt.Println("Time taken: ", end)
}

func newRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}

func addDBtoRedis(ctx context.Context, client *redis.Client, query string) error {
	address := flag.String("server", "localhost:9090", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewMoviesServiceClient(conn)

	req := v1.ReadRequest{
		Api:        apiVersion,
		MovieGenre: "Action",
	}
	res, err := c.GetMovieByGenre(ctx, &req)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	for _, m := range res.Movies {
		log.Printf("Movie: <%+v>\n", m)
	}

	args := make(map[int32]interface{})
	for i := range res.Movies {
		convert := struct {
			MovieName  string  `json:"MovieName"`
			MovieGenre string  `json:"MovieGenre"`
			Director   string  `json:"Director"`
			Rating     float32 `json:"Rating"`
		}{}
		convert.MovieName = res.Movies[i].MovieName
		convert.MovieGenre = res.Movies[i].MovieGenre
		convert.Director = res.Movies[i].Director
		convert.Rating = res.Movies[i].Rating

		byteArray, err := json.Marshal(convert)
		if err != nil {
			return err
		}
		ID := res.Movies[i].ID
		args[ID] = byteArray
		client.HSet(ctx, query, ID, args[ID])
	}
	return err
}
