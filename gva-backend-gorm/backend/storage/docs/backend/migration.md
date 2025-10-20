# Migration

## How to add new table or edit columns?

- add new table or column in `schema.hcl`
- run `make migrate.diff {name}` to generate an new migration file in `/sql`
- run `make migrate.apply` to execute or sync `/sql`

## How to add new `sql` for seeding?

- run `make migrate.new {name} to` to create an new empty migration file in `/sql`
- edit new migration file
- run `make migrate.apply` to execute or sync `/sql`

## How to migrate down | rollback?

- run `migrate.down.check` to check if script to use to rollback is current
- run `migrate.down.auto` will execute `sql down script` and remove last file in `/sql`
