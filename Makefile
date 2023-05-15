migrate-create:
	migrate create -ext sql -dir database/migration -seq $(name)
migrate-up-all:
	migrate -path ./database/migration -database "postgres://postgres:123456@localhost:5432/demo-gofiber-db?sslmode=disable" -verbose up
migrate-up:
	migrate -path ./database/migration -database "postgres://postgres:123456@localhost:5432/demo-gofiber-db?sslmode=disable" -verbose up $(step)
migrate-down:
	migrate -path ./database/migration -database "postgres://postgres:123456@localhost:5432/demo-gofiber-db?sslmode=disable" down $(step)
migrate-force:
	migrate -path ./database/migration -database "postgres://postgres:123456@localhost:5432/demo-gofiber-db?sslmode=disable" force $(ver)