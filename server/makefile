# ==============================================================================
# Main

run:
	go run ./cmd/passbook/main.go

build:
	go build ./cmd/passbook/main.go

redis:
	redis-server

test:
	go test -cover ./...

# ==============================================================================
# Go migrate postgresql

migrate_up:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/passbook?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/passbook?sslmode=disable" -verbose down

#generate:
#	mockgen -source redis_repository.go -destination mock/redis_repository.go -package mock
#	mockgen -source usecase.go -destination mock/usecase.go -package mock
	#mockgen -source pg-repository.go -destination mock/pg_repository.go -package mock
	#mockgen -source redis-repository.go -destination mock/redis_repository.go -package mock
	mockgen -source usecase.go -destination mock/usecase.go -package mock