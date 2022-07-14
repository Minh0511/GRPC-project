# GRPC-project
A CRUD project using gRPC, MySQL and Golang

# Using the project

In the root folder, go to the cmd directory and run the server:
````
cd cmd/server
server.exe -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>
````
In another terminal, go to the cmd directory and run the client
````
cd cmd/test-client
go build .
test-client.exe -server=localhost:9090
````
# Important:
The "go build ." command will create an .exe file which can be run on Windows only.
In order to run  our project on Linux, we have to do some extra steps

In the server folder, run:
````
chmod +x server
./server -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>
````

In the test-client folder, run:
````
chmod +x test-client
./test-client -server=localhost:9090
````

If you do these steps correctly, you'll see the project's working
