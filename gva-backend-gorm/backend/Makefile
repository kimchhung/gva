# Makefile for some commands.
include ./cmd/makefiles/*.mk


docker.network:
	docker network create -d bridge foxie-network


docker.up:
	docker compose -f docker-compose.yml up -d

docker.up.air:
	CUSTOM=Air docker compose -f docker-compose.yml up -d

docker.build:
	docker compose -f docker-compose.yml up -d --build

# Tidy
tidy:
	go mod tidy

# Stop application
docker.down:
	docker-compose -f docker-compose.yml down

docker.air:
	CUSTOM=Air docker-compose -f docker-compose.yml up api-service

test:
	find . -name '$(name)*_test.go' -execdir go test -v {} \;

cmd.build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/cmd/ ./cmd/...

