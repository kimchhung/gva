# How to start Backend in docker

## Requirement

- makefile : for make cli
- docker : for enviroment and air

## How ENV work ?

- `env/config.go` with goTag = `default:"val"` as template to generate `.env`
- `env/config.go` contain go types usages and validation `env` data load from `.env`
- `make env.create` will generate `.env` from `env/config.go` if file is not exist

## Run seeding work ?

- seed will always trigger on app start
- can disable by `seed.enable` in `env/config.go`,
- can disable each parts by `seed.blacklist_types` in `env/config.go`
- after changed default config in `env/config.go` delete file `.env` then run cmd `make env.create`

## How to start this project ?

- `make env.create` to create `.env` from `env/config.go`
- `make docker.network` share network for db,cache, `api-service`
- `make docker.up.air` to up db,cache,migration and hot reload on code change re-run `api-service`

## More Document

- please check in `/storage/backend`
