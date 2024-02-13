# GVA GO VUE ADMIN Boilerplate
[![Go Reference](https://pkg.go.dev/badge/api.svg)](https://pkg.go.dev/api)

Simple and scalable boilerplate to build powerful and organized REST projects with [Fiber](https://github.com/gofiber/fiber). 

## Directory Structure

```
├── app
│   ├── database
│   │   ├── schema
│   │   │   └── article.go
│   │   └── seeder
│   │       └── article_seeder.go
│   ├── middleware
│   │   ├── register.go
│   │   └── token
│   │       └── token.go
│   ├── module
│   │   └── article
│   │       ├── article_module.go
│   │       ├── controller
│   │       │   ├── article_controller.go
│   │       ├── repository
│   │       │   ├── article_repository.go
│   │       ├── dto
│   │       │   └── article_request.go
│   │       ├── service
│   │       │   └── article_service.go
│   │       └── article_module.go
│   └── router
│       └── api.go
├── build
│   ├── Dockerfile
│   └── DockerfileAir
├── cmd
│   └── example
│       ├── generate.go
│       └── main.go
├── config
│   └── example.toml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   └── bootstrap
│       ├── database
│       │   └── database.go
│       ├── logger.go
│       └── webserver.go
├── LICENSE
├── Makefile
├── README.md
├── storage
│   ├── ascii_art.txt
│   ├── private
│   │   └── example.html
│   ├── private.go
│   └── public
│       └── example.txt
└── utils
    ├── config
    │   └── config.go
    ├── response
    │   ├── response.go
    │   └── validator.go
    └── utils.go
```

## Usage
You can run that commands to run project:

```go mod download```

```go run cmd/example/main.go``` or ```air -c .air.toml``` if you want to use air

### Docker
```shell
docker-compose build
docker-compose up

CUSTOM="Air" docker-compose up # Use with Air
```

## Tech Stack
- [Go](https://go.dev)
- [Mysql](https://www.postgresql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Ent](https://github.com/ent/ent)
- [Fx](https://github.com/uber-go/fx)
- [Zerolog](https://github.com/rs/zerolog)
- [GoMock](https://github.com/golang/mock)

## To-Do List
- [x] More error-free logging.
- [x] Add makefile to make something shorter.
- [x] Introduce repository pattern.
- [ ] Add unit tests.
- [x] Add mocking with GoMock.

## License
api is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).



## migration


## migration generate
1. edit schema in "app/database/schema"
2. genrate ent using ```make gen name="add table"```
3. apply migrations ```make migrate.apply```

## migration manual
1. create migration files ```make migrate.new name="add_user_data"```
2. generate hash  ```make migrate.hash```
3. apply migrations ```make migrate.apply```

## migration role back

1. remove one or two latest migration files in "app/database/migrations"
2. check different and apply ```make migrate.schema.apply```
3. set version migration to match current version in folder ``` make migrate.apply ```