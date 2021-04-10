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

container.build: build.linux
	docker build . -t rent
container.run: container.build
	docker run --env-file=./dev/integration-tests/config.env \
          --network=rent_network \
          --name=rent \
          -p 8080:8080 \
          --rm \
          -d \
          rent