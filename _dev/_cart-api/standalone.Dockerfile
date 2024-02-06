FROM golang:latest

WORKDIR /go/src/app

COPY ./build/cart-api .

ENV PORT=${CARTAPI_SERVER_PORT}

EXPOSE $PORT

CMD ["./cart-api"]
