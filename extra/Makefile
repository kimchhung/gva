# Makefile for some commands.
include migrate.mk

# Generate Ent 
gen:
	go generate cmd/ent/generate.go

# Run both dashboad and web same port
run: 
	go run api/main.go

# Run only web api
run.web: 
	go run api/web/main.go

# run only admin
run.admin: 
	go run api/admin/main.go

swag: 
	make swag.web 
	make swag.admin

swag.web: 
	swag init -g main.go -d ./api/web --parseDependency --output ./api/web/docs/ --instanceName web

swag.admin:
	swag init -g main.go -d ./api/admin --parseDependency --output ./api/admin/docs --instanceName admin

crud:
	go run cmd/code_gen/generate.go $(name)

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

      