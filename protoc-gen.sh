#protoc --proto_path=api/proto/movie --proto_path=third_party --go_out=plugins=grpc:pkg/api/ movie-service.proto

protoc -I . -I ./googleapis \
    --go_out ./pkg \
    --go_opt paths=source_relative \
    --go-grpc_out ./pkg \
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./pkg \
    --grpc-gateway_opt paths=source_relative \
    ./api/proto/v1/movie-service.proto