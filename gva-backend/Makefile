# Makefile for some commands.
include cmd/makefiles/*.mk



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

