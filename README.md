# Wagesum demo appication
Demo microservice with GO programming language. It can calculate the sum of salaries 
under manager recursively. It supports CRUD functionality for employees and salaries.

REST interface details are available [wagesum-openapi.yaml](./api/wagesum-openapi.yaml) 

## Howto build and start

You need gmake and go 1.19 installed. (1.17+)

```shell
docker run -p 5432:5432 --name wagesum-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
make build
./out/bin/wagesum
```

## Configuration
This application can be configured runtime with environment variables.

| Env var name              | Type           | Default value    |
| ------------------------- | -------------- | ----------------:|
| WAGESUM_DB_TYPE           | string         | 127.0.0.1        |
| WAGESUM_DB_HOST           | string         | 5432             |
| WAGESUM_DB_PORT           | string         | wagesum          |
| WAGESUM_DB_NAME           | string         | 127.0.0.1        |
| WAGESUM_DB_USERNAME       | string         | postgres         |
| WAGESUM_DB_PASSWORD       | string         | mysecretpassword |
| WAGESUM_HTTP_SERVER_PORT  | string         | 3000             |

## Documentation
* [English documentation](docs/desc_en.md)
* [Hungarian documentation](docs/desc_hu.md)