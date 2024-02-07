# interceptor

_**interceptor**_, is a server-side component that intercepts incoming HTTP requests, collects relevant signals and sends a request to another API to validate whether the request was made by a human or a bot.

## Prerequisites

- go >= 1.21.0
- docker >= 20.10.23

**To play with interceptor : [Check Test With Docker section](#with-Docker)**

> **_NOTE:_**  The repo is private so if `go get github.com/dre-zouhair/interceptor` fail, you should update the `GOPRIVATE` Environment variable. (verify `GOPRIVATE` old value before running `go env -w GOPRIVATE=github.com/dre-zouhair/interceptor`).

> **_NOTE:_**  The repo is private, so building a docker image for a project that has `github.com/dre-zouhair/interceptor` as a dependency is complicated.

## How to use it?

_**interceptor**_ can be used as a :

### Standalone component

_**interceptor**_ can be deployed to act as a gateway or proxy. There are two ways of archiving this, via a **docker image** or as an executable.

> **_NOTE:_**  As an executable, you can clone the git repo and build it, or use the module as a dependency, and then register the **Interceptor** handler

### Integrated as a middleware

_**interceptor**_ be used also as a middleware inside your applications.

start by adding the module

```bash
go get github.com/dre-zouhair/interceptor
```
and then you can create the middleware instance as follows:

```text
import (
        "github.com/dre-zouhair/interceptor/middleware"
        interceptorconf "github.com/dre-zouhair/interceptor/config"
       )
....
// initialize the middleware configuration
middlewareConf := interceptorconf.ProtectionMiddlewareConfig{}

// create a protection middleware instance
protectionMiddleware := middleware.NewProtectionMiddleware(middlewareConf)

```
Currently, there are two middleware implementation available

1. net/http router
    ```text
	StandardProtectionMiddleware(next http.Handler) http.Handler
    ```

2. bunrouter middleware
    ```text
    BunProtectionMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc
    ```

## _**interceptor**_ configuration

_**interceptor**_ can be configured with the following env variables :

```dotenv
INTERCEPTOR_SERVER_PORT=#interceptor port
INTERCEPTOR_INTERCEPTION_PATH=#the path to intercept
# ALLOW|VERIFY|BLOCK
INTERCEPTOR_PROTECTION_FAIL_MODE=#the default return policy if the protection api is KO
INTERCEPTOR_PROTECTION_ENDPOINT=#the protetcion endpoint
INTERCEPTOR_PROTECTION_TOKEN=#the protetcion api authorization token
INTERCEPTOR_FORWARD_ENDPOINT=#the end point to which incoming requests are redirected after validation
# separated with ,
INTERCEPTOR_PROTECTION_CUSTOM_HEADERS=#list of headers to collect as signals for the protection api
# separated with ,
INTERCEPTOR_PROTECTION_CUSTOM_COOKIES=#list of headers to collect as cookies for the protection api
```

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

## Build **_interceptor_** executable

To build **_interceptor_** as an executable

```bash
# for dev only, this will load configuration vars form .env file
go build -tags dev ./cmd/api/main.go
```

```bash
go build ./cmd/api/main.go
```

## Run **_interceptor_** executable

To run the api execute the following command :
```bash
# for dev only, this will load configuration vars form .env file
go run -tags dev ./cmd/api/main.go
```

## Build & Run with Docker

### Build a docker image

```bash
docker build -t interceptor .
```

### Run a container

```bash
docker run --env-file ./dev.env -e INTERCEPTOR_SERVER_PORT=80 -e TAG='dev' -p 80:80 interceptor
```

> **_NOTE:_**  if you have changed the server port, be sure to update the docker run command


## Run and test the _**interceptor**_ in development mode

To test interceptor, 3 applications are present in the `_dev` folder

1. _cart-api : a simple api to protect
2. _protection-api : a simple protection api with basic signals and rules
3. _web-app : a simple web app that will make requests to _cart-api

### build & run each component

Download dependencies

```bash
make deps
```

Then run the protection api :

```bash
make run-protection-api
```

Then run the cart api (_**interceptor**_ is used as a middleware)

```bash
make run-protected-cart-api
```

Then run the web app :

```bash
make run-web-app
```

## with Docker

```bash
make run-dev
```

```bash
make clean-dev
```

Then you can visit the web app on :

```makefile
http://localhost:3000/index.html
```