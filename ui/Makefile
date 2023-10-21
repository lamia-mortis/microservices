dependencies:
	npm --prefix resources install
	go mod tidy

frontend:
	npm run --prefix resources build

server:
	go run cmd/main.go