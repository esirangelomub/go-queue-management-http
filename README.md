# API REST / QUEUE IN GO

A simple Queue Processor with Http Interface

## Installation

Clone the project [where](https://github.com/esirangelomub/go-queue-management-http).

```bash
git clone git@github.com:esirangelomub/go-queue-management-http.git
```

Init DB container

```bash
docker-compose up -d
```

## Init application

#### HTTP Server

```bash
go run main.go
```

#### Cron Jobs (not yet ready)

```bash
go run cron.go
```

## Compiling application and run (not yet ready)

```bash
go build
```

```bash
./go-api-rest
```
or
```bash
./go-api-rest.exe
```

## Access API

See the Postman collection with all CRUD examples.

## Todo
Implement tests.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)