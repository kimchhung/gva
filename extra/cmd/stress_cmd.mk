# Makefile for some commands.
# route file app/database/data/routes_data.json
# pull route from db to file
stress_cmd.hello:
	go run cmd/stress_cmd/main.go hello

