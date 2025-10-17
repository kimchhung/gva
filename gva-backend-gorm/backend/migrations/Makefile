# Makefile for some commands.

env ?= "local"
args ?=

ifeq ($(env),"local")
  CONFIG_FILE := local.connection.env
else
  CONFIG_FILE := ${env}.connection.env
endif

# Include the chosen .env file
include $(CONFIG_FILE)

migrate.inspect:
	atlas schema inspect -u "${url}"

migrate.apply.hash:
	atlas migrate apply \
		--dir "${dir}" \
		--url "${url}" \
		${args}

migrate.down.hash:
	atlas schema apply \
		--to "${dir}" \
		--url "${url}" \
		--dev-url "${devUrl}" \
		--exclude "atlas_schema_revisions"\
		${args}

migrate.down.auto:
	rm -f $(shell ls -t ./sql/*.sql | head -n 1)
	make down
	make set version=$(shell ls -t ./sql/*.sql | sed -n '2p' | grep -oE '[0-9]+')

migrate.set:
	atlas migrate set ${version} --url ${url} --dir ${dir}

migrate.status:
	atlas migrate status --url ${url} --dir ${dir}

migrate.down.check:
	atlas migrate down \
		--dir "${dir}" \
		--url "${url}" \
		--dev-url "${devUrl}" \
		--dry-run \
		${args}

migrate.hash:
	atlas migrate hash \
		--dir "${dir}" \
		${args}

migrate.diff: hash
	atlas migrate diff ${name} \
  --dir "${dir}" \
 	--to "file://schema.hcl" \
  --dev-url "${devUrl}"

