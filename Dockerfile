FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o interceptor ./cmd/api

ENV PORT=${INTERCEPTOR_SERVER_PORT}

EXPOSE $PORT

CMD ["./interceptor"]
