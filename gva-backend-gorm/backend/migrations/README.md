# Reset Migration

## Inspect DB schema

```bash
$ atlas schema inspect -u "mysql://root:password@localhost:3306/gva" > schema.hcl
```

## Create init migration

```bash
$ make migrate.create name=init
```

## Get new init migration

update schema.hcl with of schema of empty database

```bash
> schema "gva" and schema.lottery -> schema "test-db" and schema.test-db
```

then run:

```bash
$ make migrate.diff
```

then copy the output to the new migration file

### Migration roleback or reverse

1. remove one or two latest migration files in "./sql" folder
2. check different and apply by `make down`
3. set version migration to match current version in folder `make set version=XX`
4. check current version by `make status`
