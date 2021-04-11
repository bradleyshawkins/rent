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
	docker-compose -f docker-compose.yaml -f docker-compose.rent.yaml up -d
service.container.stop:
	docker-compose -f docker-compose.yaml -f docker-compose.rent.yaml down -v
service.test.integration: service.container.run test.integration service.container.stop