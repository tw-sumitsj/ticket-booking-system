# Ticket Booking System

## Architecture
![enter image description here](https://drive.google.com/uc?id=1ztd-hCJqkuw6vuRU0SWTarzalZPDQaAI)

## Dependencies
- Payment Service
- Notification Service
- Cinema Catalog Service

## Setup & Build

1. Clone the repo in the $GOPATH/src/ directory.
2. `go run main.go --migrate` to run migrations


## start
`go run main.go` to start the server

## Commands & Tips

`go run main.go --help` to see all the available options

`docker-compose exec server bookingsystem --migrate` to migrate  with docker compose
`docker-compose exec server bookingsystem --rollback` to rollback  with docker compose
`docker-compose exec server bookingsystem --force_fix_dirty` to fix any migration issue

