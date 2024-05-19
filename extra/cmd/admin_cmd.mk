# Makefile for some commands.
# route file app/database/data/routes_data.json
# pull route from db to file
admincmd.pull.route:
	go run cmd/admincmd/main.go pull.route

# push route from file to db
admincmd.push.route:
	go run cmd/admincmd/main.go push.route

# push route from file to db
admincmd.seeds:
	go run cmd/admincmd/main.go seeds

# seed all permission to db
admincmd.seed.permission:
	go run cmd/admincmd/main.go seed.permission

