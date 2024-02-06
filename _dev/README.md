# Protetcion-api

This is an API for simulating a basic protection API

## Prerequisites

- go >= 1.21.0

## Unit tests

## Build & Run

## Build & Run with Docker

### Build a docker image

```bash
docker build -t protetcion-api .
```

### Run api container

```bash
docker run -e PORT=5789 -p 5789:5789 protetcion-api
```

> **_NOTE:_**  if you have changed the server port, be sure to update the docker run command



