# DB setup

## 1 initialize db image

`docker compose up -d gva-backend-mysql`

## 2 connect docker db and create schema dev

`CREATE SCHEMA IF NOT EXISTS gva_backend_dev DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;`

## 3 run migration

`make migrate.apply`

## 4 run seeds

### Run all api modules

`make api().run`

### Run specific module

`make api().run.<module_name>`

e.g. `make api().run.admin` or `make api().run.bot`
