# ==============================================================================
# Main

run:
	go run ./cmd/passbook/main.go

build:
	go build ./cmd/passbook/main.go

test:
	go test -cover ./...


# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...


# ==============================================================================
# Go migrate postgresql

force:
	migrate -database "postgres://postgres:55555@127.0.0.1:5432/passbook?sslmode=disable" -path migrations force 1

version:
	migrate -database postgres://postgres:55555@localhost:5432/passbook?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:55555@localhost:5432/passbook?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:55555@localhost:5432/passbook?sslmode=disable -path migrations down 1

