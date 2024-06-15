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
    │   ├── request.go
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
- [Swagger](https://github.com/swaggo/swag)


## CRUD generator

```make crud name="todo" ```

- example todo CRUD
```
│── docs
│   ├── swagger.json
│   └── swagger.yaml
│── app
│   ├── database
|   |   └── schema
|   |         └── todo.go
│   ├── module
│   │   └── router.go
│   │   └── todo
│   │       ├── todo_module.go
│   │       ├── controller
│   │       │   ├── todo_controller.go
│   │       ├── repository
│   │       │   ├── todo_repository.go
│   │       ├── dto
│   │       │   └── todo_request.go
│   │       ├── service
│   │       │   └── todo_service.go
│   │       └── todo_module.go


➜  api git:(main) ✗ make crud name=todo_you
go run cmd/code_gen/generate.go todo_you
Generated app/database/schema/todo_you.go
Generated app/module/todo_you/todo_you_module.go
Generated app/module/todo_you/dto/todo_you_request.go
Generated app/module/todo_you/repository/todo_you_repository.go
Generated app/module/todo_you/service/todo_you_service.go
Generated app/module/todo_you/controller/todo_you_controller.go

2024/02/17 04:44:01 Generate swagger docs....
2024/02/17 04:44:01 Generate general API Info, search dir:./
2024/02/17 04:44:01 Generating request.Response
2024/02/17 04:44:01 Generating dto.AdminRequest
2024/02/17 04:44:01 Generating dto.TodoYouRequest
2024/02/17 04:44:01 create docs.go at docs/docs.go
2024/02/17 04:44:01 create swagger.json at docs/swagger.json
2024/02/17 04:44:01 create swagger.yaml at docs/swagger.yaml

```

## Migration

 - Install the Atlas CLI. You can find installation instructions [here](https://atlasgo.io/integrations/go-sdk).
 - Run ```make migrate.hash``` whenever got error hash mismatched
 - Check ```migrate.mk``` for more migration cli


### Migration generate

1. edit schema in "app/database/schema"
2. genrate ent using ```make migrate.gen name="add_todo_index"```
3. apply migrations ```make migrate.apply```

### Create migration manually 
1. create migration files ```make migrate.new name="add_user_data"```
2. generate hash  ```make migrate.hash```
3. apply migrations ```make migrate.apply```

### Migration roleback or reverse

1. remove one or two latest migration files in "app/database/migrations"
2. check different and apply by ```make migrate.schema.apply```
3. set version migration to match current version in folder ``` make migrate.apply ```
4. check current version by ```make migrate.status```

## Resources
 - https://github.com/efectn/fiber-boilerplate 


## License
api is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).
