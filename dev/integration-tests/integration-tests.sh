#!/bin/sh

docker build . -t rent:local
docker-compose up -d --force-recreate
#docker run --env-file=./dev/integration-tests/config.env \
#      --network=rent_network \
#      --name=rent \
#      --rm \
#      -p 8080:8080 \
#      -d \
#      rent

 go test -tags=integration ./...

 echo "finished tests"

# docker kill rent
# docker rmi rent
#
# docker-compose down -v