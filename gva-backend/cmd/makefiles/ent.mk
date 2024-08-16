# Generate Ent
ent.gen:
	go generate cmd/ent/generate.go

# Generate EntGql
ent.gen.gql:
	go generate cmd/ent/generategql.go