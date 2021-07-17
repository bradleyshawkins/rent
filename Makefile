deps.up:
	docker-compose up -d
deps.down:
	docker-compose down -v

test.container.build:
	docker build . -f dev/integration-tests/test.Dockerfile -t rent-test:local
test.container.down:
	docker-compose -f docker-compose.yaml -f dev/integration-tests/docker-compose.test.yaml down -v
test.container.integration: test.container.down service.container.build test.container.build
	docker-compose -f docker-compose.yaml -f dev/integration-tests/docker-compose.test.yaml up --abort-on-container-exit --force-recreate -V
	docker-compose -f docker-compose.yaml -f dev/integration-tests/docker-compose.test.yaml down -v
test.unit:
	go test -tags=unit ./...

service.container.build:
	docker build . -t rent:local
service.container.run: service.container.build
	docker-compose -f ./dev/integration-tests/docker-compose.yaml up -d
service.container.stop:
	docker-compose -f ./dev/integration-tests/docker-compose.yaml down -v