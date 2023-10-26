ifndef GATEWAY_BINARY_FILE
	GATEWAY_BINARY_FILE=gateway
endif

SERVER_CMD_FILE=cmd/main.go

protobuf:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

server_run:
	go mod tidy
	go run ${SERVER_CMD_FILE}

server_build:
	CGO_ENABLED=0 go build -o ${GATEWAY_BINARY_FILE} ${SERVER_CMD_FILE}