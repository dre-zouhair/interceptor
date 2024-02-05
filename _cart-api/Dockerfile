FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o cart-api ./cmd/api

EXPOSE 5050

CMD ["./cart-api"]
