# Sneaky Spider Web App

## API Documentation

https://documenter.getpostman.com/view/17475694/2s93CPrCw6

## Getting started (Setup & Running)

### Prerequisites

Install the [Go programming language](https://go.dev/).

Install [docker](https://www.docker.com/). 
<!-- and [docker-compose](https://docs.docker.com/compose/install/). -->

Then install Go modules

```bash
go mod tidy
```

### Running the application

<!-- To start up the Postgres instance.

```bash
docker compose up -d
``` -->

To run database migrations. "D:/Projects/go" should be replaced with your local path to the project.

```bash
docker run -v D:/Projects/sneaky-spider-api/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "<your-postgres-address>" up
```

To clear the database:

```bash
docker run -v D:/Projects/sneaky-spider-api/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "<your-postgres-address>" down -all
```

To start the API:

```bash
make run
``` 
or

```bash
go run main.go
```

To start the Alerts Service:

```bash
go run alerts/cmd/main.go
```

<!-- To start the Encryption Tool:

```bash
go run encryptor/main.go
``` -->

## Testing

```bash
make test
```
or 

```bash
go test -v ./...
```

<!-- 
## Other

To make queries to the Docker Postgres instance.

```bash
docker exec -it postgres psql -U postgres -d af1spider -c "<sql query>"

``` -->
