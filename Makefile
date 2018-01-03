setup:
	docker-compose build

db:
	docker-compose up -d db
	docker-compose up -d redis

test:
	make db
	go test ./...

start:
	make test
	docker-compose up multilogin

fmt:
	go fmt ./...