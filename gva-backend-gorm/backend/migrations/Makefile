# Makefile for some commands.
env=local

migrate.inspect:
	atlas schema inspect --env ${env}

migrate.set: migrate.hash
	atlas migrate set ${version} --env ${env}

migrate.status:
	atlas migrate status --env ${env}

migrate.hash:
	atlas migrate hash --env ${env}

migrate.new:
	atlas migrate new --env ${env}

migrate.diff: migrate.hash
	atlas migrate diff ${name} --env ${env}

migrate.apply: migrate.hash
	atlas migrate apply --env ${env}

# check sql remove query
migrate.down.check:
	atlas migrate down --env ${env} --dry-run

# trigger sql remove query
migrate.down: migrate.hash
	atlas migrate down --env ${env}

# 1- run query, 2-remove file locally, 3- set current local version
migrate.down.auto:
	make migrate.down
	rm -f $(shell ls -t ./sql/*.sql | head -n 1)
	make migrate.set version=$(shell ls -t ./sql/*.sql | sed -n '2p' | grep -oE '[0-9]+')

