# Makefile for some commands.
# route file app/database/data/menu_data.json
# pull route from db to file
admincmd.pull.menu:
	go run cmd/admincmd/main.go pull.menu

# push route from file to db
admincmd.push.menu:
	go run cmd/admincmd/main.go push.menu

# push route from file to db
admincmd.seeds:
	go run cmd/admincmd/main.go seeds

# seed all permission to db
admincmd.seed.permission:
	go run cmd/admincmd/main.go seed.permission

.PHONY: admincmd.gen

admincmd.gen:
	@echo "Running admincmd.gen with name: $(name), option: $(option)"
	@go run cmd/admincmd/main.go crud "$(name)" "$(option)"
