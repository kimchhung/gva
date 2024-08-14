# Makefile for some commands.
include cmd/migrate.mk
include cmd/admin_cmd.mk
include cmd/serve_cmd.mk
include cmd/swag_cmd.mk

# Generate Ent
gen:
	go generate cmd/ent/generate.go

# Generate EntGql
gen.gql:
	go generate cmd/ent/generategql.go


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

test:
	find . -name '$(name)*_test.go' -execdir go test -v {} \;

