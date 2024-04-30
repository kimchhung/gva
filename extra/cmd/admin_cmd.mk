# Makefile for some commands.
# route file app/database/data/routes_data.json
# pull route from db to file
admin_cmd.pull.route:
	go run cmd/admin_cmd/main.go pull.route

# push route from file to db
admin_cmd.push.route:
	go run cmd/admin_cmd/main.go push.route

# push route from file to db
admin_cmd.seeds:
	go run cmd/admin_cmd/main.go seeds

# seed all permission to db
admin_cmd.seed.permission:
	go run cmd/admin_cmd/main.go seed.permission

