reset-db:
	sh ./src/script/reset_db.sh 

migrate:
	go run src/script/migrate.go

run:
	go run src/main.go -config ./config.toml

build:
	go build -o ./bin/main ./src/main.go 

dev-run:
	go build -o ./bin/main ./src/main.go
	./bin/main
