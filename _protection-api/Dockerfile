FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o protection-api ./cmd/api

ENV PORT=5789

EXPOSE $PORT

CMD ["./protection-api"]
