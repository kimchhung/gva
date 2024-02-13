# GVA GO VUE ADMIN Boilerplate
[![Go Reference](https://pkg.go.dev/badge/api.svg)](https://pkg.go.dev/api)

Simple and scalable boilerplate to build powerful and organized REST projects with [Fiber](https://github.com/gofiber/fiber). 

## Directory Structure

```
├── app
│   ├── database
│   │   ├── migrations 
│   │   ├── schema
│   │   │   └── article.go
│   │   ├── seeder
│   │   │    └── article_seeder.go
│   │   └── local.connection.env  // migration connection
│   ├── middleware
│   │   ├── register.go
│   │   └── token
│   │       └── token.go
│   ├── module
│   │   └── router.go
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
├── build
│   ├── Dockerfile
│   └── DockerfileAir
├── cmd
│   └── code_gen
│   └── ent
│       ├── generate.go
│       └── migrate
├── config
│   │── config.go
│   └── config.toml
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
├── migrate.mk
├── README.md
├── tools.go
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

```go run main.go``` or ```air -c .air.toml``` if you want to use air

### Docker
```shell
docker-compose build
docker-compose up

CUSTOM="Air" docker-compose up # Use with Air
```

## Tech Stack
- [Go](https://go.dev)
- [Mysql](https://www.mysql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Ent](https://github.com/ent/ent)
- [Atlas](https://atlasgo.io)
- [Fx](https://github.com/uber-go/fx)
- [Zerolog](https://github.com/rs/zerolog)
- [CodeGen](https://github.com/dolmen-go/codegen)

## License
api is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).



## Migration

 - Install the Atlas CLI. You can find installation instructions here.
run ```make migrate.hash``` whenever got error hash mismatched


### Migration generate

1. edit schema in "app/database/schema"
2. genrate ent using ```make gen name="add table"```
3. apply migrations ```make migrate.apply```

### Create migration manual 
1. create migration files ```make migrate.new name="add_user_data"```
2. generate hash  ```make migrate.hash```
3. apply migrations ```make migrate.apply```

### Migration role back

1. remove one or two latest migration files in "app/database/migrations"
2. check different and apply ```make migrate.schema.apply```
3. set version migration to match current version in folder ``` make migrate.apply ```


## Resources
 - https://github.com/efectn/fiber-boilerplate