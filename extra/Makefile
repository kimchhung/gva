# Makefile for some commands.
include migrate.mk

# Generate Ent 
gen:
	go generate cmd/ent/generate.go

run.web: 
	go run api/web/web.go

run.dash: 
	go run api/dashboard/dashboard.go

swag.web:
	swag init -g web.go -d ./api/web --parseDependency --output ./docs/web/ --instanceName web

swag.dash:
	swag init -g dashboard.go -d ./api/dashboard --parseDependency --output ./docs/dashboard/ --instanceName dashboard

crud:
	go run cmd/code_gen/generate.go $(name)

# Run application outside Docker
run:
	go run main.go

# Run application inside Docker
docker:
	docker-compose up

# Run application inside Docker with Air
air:
	CUSTOM="Air" docker-compose up --build example

# Tidy
tidy:
	go mod tidy

# Stop application
stop:
	docker-compose down

