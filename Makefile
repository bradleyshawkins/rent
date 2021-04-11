deps.up:
	docker-compose up -d
deps.down:
	docker-compose down -v

test.integration:
	./dev/integration-tests/integration-tests.sh

build.linux:
	./dev/build.sh
build.mac:
	GOOS=darwin ./dev/build.sh

container.build:
	docker build . -t rent:latest --no-cache
container.run: deps.up container.build
	docker run --env-file=./dev/integration-tests/config.env \
          --network=rent_network \
          --name=rent \
          --rm \
          -p 8080:8080 \
          -d \
          rent:latest
container.stop:
	docker rm -f rent