migrate-up:
	migrate -path=sql/migrations -database "postgres://root:root@tcp(localhost:5432)/mydb" -verbose up

migrate-down:
	migrate -path=sql/migrations -database "postgres://root:root@tcp(localhost:5432)/mydb" -verbose down

docker-up:
        docker-compose up -d

docker-down:
        docker-compose down

sqlc-generate:
        sqlc generate

dev-setup: docker-up
        @echo "Aguardando PostgreSQL iniciar..."
        @sleep 5
        make migrate-up
        make sqlc-generate	

dev-start: dev-setup
        go run cmd/main.go

.PHONY: docker-up docker-down sqlc-generate dev-setup dev-start migrate-up migrate-down