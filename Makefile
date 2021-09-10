deps.up:
	docker-compose up -d
deps.down:
	docker-compose down -v

test.unit:
	go test -tags=unit ./...
test.integration:
	go test -tags=integration ./...
test.integration.env:
	./dev/integration.sh


service.container.build:
	docker build . -t rent:local
service.container.run: service.container.build
	docker-compose -f docker-compose.yaml up -d

env.start: service.container.build
	docker-compose -f docker-compose.yaml -f dev/docker-compose.local.yaml up -d --force-recreate -V
env.stop:
	docker-compose -f docker-compose.yaml -f dev/docker-compose.local.yaml down