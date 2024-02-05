FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get ./...

RUN go build -o interceptor ./cmd/api

ENV PORT=7777

EXPOSE $PORT

CMD ["./interceptor"]
