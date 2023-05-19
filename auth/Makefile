postgres:
	docker run --name work_room_auth_postgres -e POSTGRES_PASSWORD=password  -p 5432:5432 --rm -d postgres

redis:
	docker run --name work_room_auth_redis -p 6379:6379 --rm -d redis


create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migrations -seq {migration_name}

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations down