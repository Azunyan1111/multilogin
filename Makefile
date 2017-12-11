test:
	docker-compose up -d db
	go test ./...

start:
	make test
	docker-compose up

fmt:
	go fmt ./...