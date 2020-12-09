# Go Boilerplate

## Installation

1. `pip install pre-commit`
2. `go mod download`
3. install [docker](docs.docker.com)
4. setup `.env` from `.env.sample`

## Developing

1. run `docker-compose up -d`
2. run migration `go run . migrate`
3. run server `go run . serve`
