gen-protoc:
	protoc --go_out=. --go-grpc_out=. proto/*.proto  --experimental_allow_proto3_optional

postgres:
	docker run --name work_room_projects_tasks_postgres -e POSTGRES_PASSWORD=password  -p 5432:5432 --rm -d postgres


create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migrations -seq init

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations down

migrate-drop:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations drop