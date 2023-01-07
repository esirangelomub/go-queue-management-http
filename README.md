# API REST IN GO

A simple Rest API made in Go

## Installation

Clone the project [where](git@github.com:esirangelomub/go-api-rest.git).

```bash
git clone git@github.com:esirangelomub/go-api-rest.git
```

Init the container

```bash
docker-compose up -d
```

## Init application

#### HTTP Server

```bash
go run server.go
```

#### Cron Jobs

```bash
go run cron.go
```

## Compiling application and run

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