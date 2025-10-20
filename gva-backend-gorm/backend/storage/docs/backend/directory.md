# Backend Directory Structure

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
│   └── main.go
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
