# interceptor

This is an API for simulating adding items to a shopping cart.

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

```bash
go build -tags dev ./cmd/api/main.go
```

## Run the api

To run the api execute the following command :
```bash
go run -tags dev ./...
```

## Build & Run with Docker

### Build a docker image

```bash
docker build -t interceptor .
```

### Run api container

```bash
docker run -e PORT=80 -e TAG='dev' -p 80:80 interceptor
```

> **_NOTE:_**  if you have changed the server port, be sure to update the docker run command



