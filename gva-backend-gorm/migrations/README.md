# Reset Migration

## Inspect DB schema
```bash
$ atlas schema inspect -u "mysql://user:123456@localhost:3306/lottery" > schema.hcl
```

## Create init migration
```bash
$ make create name=init
```

## Get new init migration
update schema.hcl with of schema of empty database

```bash
> schema "lottery" and schema.lottery -> schema "test-db" and schema.test-db
```

then run:

```bash
$ make diff
```

then copy the output to the new migration file
