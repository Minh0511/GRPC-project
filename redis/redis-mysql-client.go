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

	searchMovieGenre := "Horror"

	nameInRedis, err := redisClient.HGetAll(ctx, searchMovieGenre).Result()

	if err != nil {
		log.Println(err)
		panic(err)
	}
	if len(nameInRedis) == 0 {
		log.Println("No movie found")
		res := addDBtoRedis(ctx, redisClient, searchMovieGenre)
		if res != nil {
			log.Println(res)
			panic(res)
		}
	} else {
		start := time.Now()
		log.Println("Found movie in redis")
		for key, value := range nameInRedis {
			log.Printf("Movie: <%+v>\n", key)
			log.Printf("Movie: <%+v>\n", value)
		}
		log.Println("Time taken: ", time.Since(start))
	}
	result, err := redisPing(redisClient, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func newRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}

func redisPing(client *redis.Client, ctx context.Context) (string, error) {
	var result, err = client.Ping(ctx).Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
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
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	req := v1.ReadRequest{
		Api:        apiVersion,
		MovieGenre: "Action",
	}
	res, err := c.GetMovieByGenre(ctx, &req)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	fmt.Println("Get movies by genre:")
	for _, m := range res.Movies {
		log.Printf("Movie: <%+v>\n", m)
	}

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
		client.HSet(ctx, query, res.Movies[i].MovieName, byteArray)
	}
	return err
}
