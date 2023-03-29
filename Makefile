run:
	go run cmd/birth_minder/main.go

migrate:
	go run cmd/birth_minder/main.go -migration

build:
	go mod download && CGO_ENABLED=0 go build -o ./bin/bminder_app ./cmd/birth_minder/main.go