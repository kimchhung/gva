# Makefile for some commands.
# route file app/database/data/routes_data.json
# pull route from db to file
admin_cmd.pull:
	go run cmd/admin_cmd/main.go pull

# push route from file to db
admin_cmd.push:
	go run cmd/admin_cmd/main.go push

# push route from file to db
admin_cmd.seeds:
	go run cmd/admin_cmd/main.go seeds



