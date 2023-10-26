ifndef FRONTEND_SRC_DIR
	FRONTEND_SRC_DIR=resources
endif

ifndef SERVER_BINARY_FILE
	SERVER_BINARY_FILE=staticServer
endif

SERVER_CMD_FILE=cmd/main.go


frontend:
	npm --prefix ${FRONTEND_SRC_DIR} install
	npm run --prefix ${FRONTEND_SRC_DIR} build

server_run:
	go mod tidy
	go run ${SERVER_CMD_FILE}

server_build:
	CGO_ENABLED=0 go build -o ${SERVER_BINARY_FILE} ${SERVER_CMD_FILE}