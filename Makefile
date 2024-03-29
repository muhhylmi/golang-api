# Work in windows command
# Please adjust it if use other operation system
DB_SCHEMA := $(shell powershell.exe -Command "(Get-Content .env | Select-String 'DB_SCHEMA=').ToString().Split('=')[1]")
DB_URI := $(shell powershell.exe -Command "(Get-Content .env | Select-String 'DB_URI=').ToString().Split('=')[1]")
SSL_MODE := $(shell powershell.exe -Command "(Get-Content .env | Select-String 'SSLMODE=').ToString().Split('=')[1]")
SCHEMA_NAME := $(shell powershell.exe -Command "(Get-Content .env | Select-String 'DB_POSTGRES_SCHEMA=').ToString().Split('=')[1]")

migration:
	migrate create -ext sql -dir migrations ${NAME}

migrationup:
	migrate -database "${DB_URI}?sslmode=${SSL_MODE}&search_path=${SCHEMA_NAME}" -path "${DB_SCHEMA}" up

migrationdown:
	migrate -database "${DB_URI}?sslmode=${SSL_MODE}&search_path=${SCHEMA_NAME}" -path ${DB_SCHEMA} down

migrationversion:
	migrate -database "${DB_URI}?sslmode=${SSL_MODE}&search_path=${SCHEMA_NAME}" -path ${DB_SCHEMA} version

migrationchange:
	migrate -database "${DB_URI}?sslmode=${SSL_MODE}&search_path=${SCHEMA_NAME}" -path ${DB_SCHEMA} force ${VERSION}

run:
	go run ./...

dev:
	air server

install:
	go mod download

tidy:
	go mod tidy

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

mock:
	mockery --all --recursive=true --keeptree

test:
	go test ./...