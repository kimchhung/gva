
env=local

# push route from file to db
admincmd.seed.all:
	go run cmd/admincmd/main.go seed.all ${env}

# seed all permission to db

admincmd.seed.permission:
	go run cmd/admincmd/main.go seed.permission ${env}

.PHONY: admincmd.gen

admincmd.gen:
	@echo "Running admincmd.gen with name: $(name), option: $(option)"
	@go run cmd/admincmd/main.go gen "$(name)" "$(option)"
