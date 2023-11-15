dockerinit:
	docker run --name=todo-db -e POSTGRES_PASSWORD="qwerty" -p 5434:5432 -d postgres
run:
	go run cmd/main.go

.PHONY: run dockerinit
