# Makefile for some commands.
include migrate.mk

# Generate Ent 
gen:
	go generate cmd/ent/generate.go

swag:
	swag init

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

