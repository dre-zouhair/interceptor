# Protetcion-api

This is an API for simulating a basic protection API

## Prerequisites

- go >= 1.21.0

## Unit tests
### Test with coverage
```bash
go test -cover ./...
```

### Test with coverage profile output

```bash
go test ./... -coverprofile coverage.out
```

to display the coverage output as HTML

```bash
go tool cover -html coverage.out -o coverage.html
```

## Build the api

### Linux
```bash
go build -o ./build/protetcion-api ./cmd/api/main.go
```

### Windows

```bash
go build -o .\build\protetcion-api.exe ./cmd/api/main.go
```

## Run the api

To run the api execute the following command :
```bash
go run ./...
```

To ping the api :

```bash
curl http://localhost:5789/api/v1/ping
```

## Build & Run with Docker

### Build a docker image

```bash
docker build -t protetcion-api .
```

### Run api container

```bash
docker run -e PORT=5789 -p 5789:5789 protetcion-api
```

> **_NOTE:_**  if you have changed the server port, be sure to update the docker run command



