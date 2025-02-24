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

migrate.apply:
	atlas migrate apply \
		--dir "${dir}" \
		--url "${url}" \
		${args}

migrate.hash:
	atlas migrate apply \
		--dir "${dir}" \
		--url "${url}" \
		${args}

migrate.status:
	atlas migrate status \
		--dir "${dir}" \
		--url "${url}"

migrate.diff:
	atlas migrate diff ${name} \
  --dir "${dir}" \
 	--to "file://schema.hcl" \
  --dev-url "${devUrl}"

