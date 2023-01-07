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

## Create database

```bash
docker exec -it api-rest-db-1 psql -U postgres 
```
```sql
postgres=# CREATE DATABASE api_rest WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';
```
```sql
postgres=# postgres=# CREATE USER user_rest WITH encrypted password '123456';
```
```sql
postgres=# postgres=# CREATE USER user_rest WITH encrypted password '123456';
```
```sql
postgres=# \c api_rest;
```
```sql
api_rest=# CREATE TABLE public.todos (
      id serial primary key ,
      title character varying(255),
      description text,
      done bool
);
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user_rest;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO user_rest;
```

## Init application

```bash
go run main.go
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