FROM alpine:latest

WORKDIR /app

COPY ./build/protected-cart-api  /app/

RUN chmod +x /app/protected-cart-api


ENV PORT=${CARTAPI_SERVER_PORT}

EXPOSE $PORT

CMD ["./protected-cart-api"]
