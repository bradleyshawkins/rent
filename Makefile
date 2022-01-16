deps.up:
	docker-compose up -d
deps.down:
	docker-compose down -v

test.unit:
	go test ./... -v -short
test.integration: deps.down
	./dev/integration.sh
