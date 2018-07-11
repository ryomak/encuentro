mysql -uroot < ./src/script/reset_db.sql
go run ./src/script/migrate.go
