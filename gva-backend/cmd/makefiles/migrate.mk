# Makefile for some commands.

env ?= "local"
args ?=

ifeq ($(env),"local")
  CONFIG_FILE := app/database/local.connection.env
else
  CONFIG_FILE := app/database/${env}.connection.env
endif

# Include the chosen .env file
include $(CONFIG_FILE)

migrate.diff: migrate.hash
	atlas migrate diff ${name} \
	--dir ${dir} \
	--to "ent://app/database/schema" \
	--dev-url ${devUrl}

# generate migration sql diff from schema and db
migrate.lint:
	atlas migrate lint --dev-url ${devUrl} --dir ${dir} --latest 1

# make migrate.set version = ****
migrate.set:
	atlas migrate set ${version} --url ${url} --dir ${dir}

# apply migration to db, recommend to use migrate.status before apply
migrate.apply:
	atlas migrate apply --url ${url} --dir ${dir} ${args}

migrate.schema.apply:
	atlas schema  apply --url ${url} --to ${dir} --dev-url ${devUrl} --exclude "atlas_schema_revisions"

migrate.down:
	migrate down --url ${url} --dir ${dir} --dev-url ${devUrl}

migrate.status:
	atlas migrate status --url ${url} --dir ${dir}

# generate migration sql diff from schema and db
migrate.gen:
	go run -mod=mod cmd/ent/migrate/main.go ${name} ${env}

# check If the migration directory and sum file are out-of-sync
migrate.validate :
	atlas migrate validate --dir ${dir}

# make migrate.new name = "add_migration_data"
migrate.new:
	atlas migrate new ${name} --dir ${dir}

# to re-hash the contents and resolve the error, prevent concurrent migrate
migrate.hash:
	atlas migrate hash --dir ${dir} ${args}

migrate.clean:
	atlas schema clean -u ${url}



