# GVA GO VUE ADMIN Boilerplate

[![Go Reference](<https://pkg.go.dev/badge/api().svg>)](https://pkg.go.dev/api)

Simple and scalable boilerplate to build powerful and organized REST projects with [Fiber](https://github.com/gofiber/fiber).

## Backend Directory Structure

```
├── api
│   ├── admin
│   │   ├── docs
│   │   ├── main.go
│   │   └── module
│   │       └── todo
│   │           ├── dto
│   │           │   ├── todo_request.go
│   │           │   └── todo_response.go
│   │           ├── todo_controller.go
│   │           ├── todo_module.go
│   │           └── todo_service.go
│   ├── bot
│   │   ├── docs
│   │   │   ├── bot_docs.go
│   │   │   ├── bot_swagger.json
│   │   │   └── bot_swagger.yaml
│   │   ├── main.go
│   │   └── module
│   │       ├── comic
│   │       │   ├── comic_module.go
│   │       │   └── comic_service.go
│   │       ├── index
│   │       │   ├── index_controller.go
│   │       │   ├── index_module.go
│   │       │   └── index_service.go
│   │       ├── module.go
│   │       └── router.go
│   ├── main.go
│   └── web
│       ├── docs
│       │   ├── web_docs.go
│       │   ├── web_swagger.json
│       │   └── web_swagger.yaml
│       ├── graph
│       │   ├── generated
│       │   ├── model
│       │   ├── resolver
│       │   ├── schema
│       │   └── server.go
│       ├── main.go
│       └── module
│           ├── index
│           │   ├── index_controller.go
│           │   ├── index_module.go
│           │   └── index_service.go
│           ├── module.go
│           └── router.go
├── app
│   ├── app.go
│   ├── common
│   │   ├── common_module.go
│   │   ├── context
│   │   ├── controller
│   │   ├── error
│   │   ├── permission
│   │   ├── repository
│   │   └── service
│   │       └── jwt_service.go
│   ├── database
│   │   ├── create_db.sql
│   │   ├── data
│   │   ├── migrations
│   │   ├── schema
│   │   └── seeds
│   ├── middleware
│   └── router
├── build
├── cmd
├── env
├── go.mod
├── go.sum
├── gqlgen.yml
├── internal
├── lang
├── repository
├── tools.go
└── utils
```

## Usage

You can run that commands to run project:

`go mod download`

`go run main.go` or `air -c .air.toml` if you want to use air

### Docker

```shell
docker-compose build
docker-compose up

CUSTOM="Air" docker-compose up # Use with Air
```

## CRUD generator

`make admincmd.gen name="todo" `

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


➜  backend git:(main) ✗ make admincmd.gen name=todo

Generated api/admin/module/todo/todo_module.go
Generated app/common/repository/todo_repository.go
Generated app/common/permission/todo_permission.go
Generated api/admin/module/todo/dto/todo_request.go
Generated api/admin/module/todo/dto/todo_response.go
Generated api/admin/module/todo//todo_service.go
Generated app/database/schema/todo.go

```

## Migration

- Install the Atlas CLI. You can find installation instructions [here](https://atlasgo.io/integrations/go-sdk).
- Run `make migrate.hash` whenever got error hash mismatched
- Check `migrate.mk` for more migration cli

### Migration generate

1. edit schema in "app/database/schema"
2. genrate ent using `make migrate.gen name="add_todo_index"`
3. apply migrations `make migrate.apply`

### Create migration manually

1. create migration files `make migrate.new name="add_user_data"`
2. generate hash `make migrate.hash`
3. apply migrations `make migrate.apply`

### Migration roleback or reverse

1. remove one or two latest migration files in "app/database/migrations"
2. check different and apply by `make migrate.schema.apply`
3. set version migration to match current version in folder `make migrate.apply`
4. check current version by `make migrate.status`

## Resources

- https://github.com/efectn/fiber-boilerplate

## License

api is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).
