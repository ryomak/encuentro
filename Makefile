migrate:
	go run src/script/migrate.go

run:
	./bin/main -config ./config.toml

build:
	go build -o ./bin/main ./src/main.go 

dev-run:
	go build -o ./bin/main ./src/main.go
	./bin/main
