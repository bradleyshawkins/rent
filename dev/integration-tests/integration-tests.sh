#!/bin/sh
docker-compose up -d --force-recreate

docker build . -t rent
docker run --env-file=./dev/integration-tests/config.env \
      --network=rent_network \
      --name=rent \
      -p 8080:8080 \
      --rm \
      -d \
      rent

 go test -tags=integration ./...

 docker kill rent
 docker rmi rent

 docker-compose down -v