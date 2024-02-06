# Protetcion-api

This is an API for simulating a basic protection API

## Prerequisites

- go >= 1.21.0

## Unit tests

```bahs
make tests
```

## Build & Run 

Download dependencies

```bash
make deps
```

Then run the protection api :

```bash
make run-protection-api
```

Then run the cart api :

```bash
make run-protected-cart-api
```

Then run the web app :

```bash
make run-web-app
```

## Build & Run with Docker

**This is not working because the repo is private**

```bash
make run-container
```
to clean the services

```bash
make clean-container
```


