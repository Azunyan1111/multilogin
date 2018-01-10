setup:
	docker-compose build

db:
	docker-compose up -d db
	docker-compose up -d redis

test:
	make db
	make fmt
	go get ./...
	go test ./...

start:
	make db
	docker-compose up multilogin

fmt:
	go fmt ./...