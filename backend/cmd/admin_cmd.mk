# Makefile for some commands.
# route file app/database/data/menu_data.json
# pull route from db to file
admincmd.pull.Menu:
	go run cmd/admincmd/main.go pull.Menu

# push route from file to db
admincmd.push.Menu:
	go run cmd/admincmd/main.go push.Menu

# push route from file to db
admincmd.seeds:
	go run cmd/admincmd/main.go seeds

# seed all permission to db
admincmd.seed.permission:
	go run cmd/admincmd/main.go seed.permission

