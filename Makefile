setup:
	docker-compose build

db:
	docker-compose up -d db

test:
	make db
	go test ./...

start:
	make test
	docker-compose up

fmt:
	go fmt ./...