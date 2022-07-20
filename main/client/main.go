package main

import (
	"GRPC-project/config"
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

	c := v1.NewTransactionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		Customer: &v1.Customer{
			CustomerName: config.CustomerName[rand.Intn(len(config.CustomerName))],
			Phone:        config.Phone[rand.Intn(len(config.Phone))] + "-" + strconv.Itoa(rand.Intn(999-100)+100) + "-" + strconv.Itoa(rand.Intn(9999-1000)+1000),
			Email:        strconv.Itoa(rand.Intn(9999-1000)+1000) + config.Email[rand.Intn(len(config.Email))],
			Product:      config.Product[rand.Intn(len(config.Product))],
		},
	}
	res1, err := c.CreateCustomer(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	//Get movies by genre
	req2 := v1.ReadRequest{
		Api:     apiVersion,
		Product: "Găng tay",
	}
	res2, err := c.GetCustomerByProduct(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	fmt.Println("Get movies by genre:")
	for _, m := range res2.Customer {
		log.Printf("Movie: <%+v>\n", m)
	}

	//Get all movies
	req3 := v1.ReadAllRequest{
		Api: apiVersion,
	}
	res3, err := c.GetAllCustomer(ctx, &req3)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	//log.Printf("Read result: <%+v>\n\n", res3)
	fmt.Println("Get all movies:")
	for _, m := range res3.Customer {
		log.Printf("Movie: <%+v>\n", m)
	}

	// Update
	req4 := v1.UpdateRequest{
		Api:          apiVersion,
		CustomerName: config.CustomerName[rand.Intn(len(config.CustomerName))],
	}
	res4, err := c.UpdateCustomer(ctx, &req4)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update successfull: <%+v>\n\n", res4)

	// Delete
	req5 := v1.DeleteRequest{
		Api:     apiVersion,
		Product: "Vòng cổ",
	}
	res5, err := c.DeleteCustomer(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete successfull: <%+v>\n\n", res5)
}
