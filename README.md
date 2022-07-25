# GRPC-project
A CRUD project using gRPC, MySQL and Golang

# Using the project

In the root folder, run docker compose:

```
docker-compose up
```
This will start a docker container for the project.

To start the server, go to the main folder, server folder and run:

```
go run main.go
```

To start the client, go to the main folder, client folder and run:

```
go run main.go
```

To use redis, go to the main folder, redis folder and run:

```
go run redis.go
```

To run elasticsearch, go to the main folder, elasticsearch folder and run:

```
go run elastic.go
```

If you do these steps correctly, you'll see the project's working
