FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o cart-api ./cmd/api

ENV PORT=${CARTAPI_SERVER_PORT}

EXPOSE $PORT

CMD ["./cart-api"]
