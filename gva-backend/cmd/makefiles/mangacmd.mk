
# search by name
manganato.search:
	go run cmd/mangacmd/main.go manganato.search name=${name}

manganato.detail:
	go run cmd/mangacmd/main.go manganato.detail id=${id}

manganato.latest:
	go run cmd/mangacmd/main.go manganato.latest
