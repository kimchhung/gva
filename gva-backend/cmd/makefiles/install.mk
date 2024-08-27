install.air:
	curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

install.atlas:
	curl -sSf https://atlasgo.sh | sh

install.swag:
	go install github.com/swaggo/swag/cmd/swag@latest

install.all: install.air install.atlas install.swag
