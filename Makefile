deps.up:
	docker-compose up -d
deps.down:
	docker-compose down -v

test.integration:
	./dev/integration-tests/integration-tests.sh
test.unit:
	go test -tags=unit ./...

service.container.build:
	docker build . -t rent:latest
service.container.run: service.container.build
	docker-compose -f ./dev/integration-tests/docker-compose.yaml up -d
service.container.stop:
	docker-compose -f ./dev/integration-tests/docker-compose.yaml down -v
service.test.integration: service.container.run test.integration