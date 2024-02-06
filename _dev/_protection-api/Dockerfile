FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o protection-api ./cmd/api

ENV PORT=${PROTECTIONAPI_SERVER_PORT}

EXPOSE $PORT

CMD ["./protection-api"]
