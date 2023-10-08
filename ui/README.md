# WEB INTERFACE MICROSERVICE

### MANUAL SETUP

All the directives should be run from the microservice root folder

```
    npm --prefix resources install
    npm run --prefix resources build
    go mod tidy
    go run cmd/web/main.go
```