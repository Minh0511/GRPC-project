protoc -I . -I ./googleapis \
    --go_out ../movie \
    --go_opt paths=source_relative \
    --go-grpc_out ../movie \
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ../movie \
    --grpc-gateway_opt paths=source_relative \
    moverService.proto